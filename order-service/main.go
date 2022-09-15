package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/version", getVersion)
	router.GET("/orders", getOrders)
	router.GET("/orders/:orderId", getOrder)
	router.POST("/orders", postOrder)
	router.Run("0.0.0.0:8081")
}

func getVersion(c *gin.Context) {
	response := map[string]string{
		"version": os.Getenv("SERVICE_VERSION"),
	}
	c.IndentedJSON(http.StatusOK, response)
}

type invoice struct {
	InvoiceID string  `json:"invoice_id"`
	Amount    float32 `json:"amount"`
}

type order struct {
	ID      string   `json:"id"`
	Date    string   `json:"date"`
	Total   float32  `json:"total"`
	Invoice *invoice `json:"invoice"`
}

var orders = []order{
	{ID: "1", Date: "01-02-2022", Total: 1000.00},
	{ID: "2", Date: "02-02-2022", Total: 1500.00},
	{ID: "3", Date: "03-02-2022", Total: 2000.00},
}

func getOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)
}

func getOrder(c *gin.Context) {
	orderFound := false
	orderId := c.Param("orderId")
	for _, order := range orders {
		if orderId == order.ID {
			orderFound = true
			invoice, err := getInvoice(orderId)
			if err != nil {
				panic(err)
			} else {
				order.Invoice = invoice
			}
			c.IndentedJSON(http.StatusOK, order)
		}
	}
	if !orderFound {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getInvoice(orderId string) (*invoice, error) {
	resp, err := http.Get(fmt.Sprintf("http://invoice-service:8082/orders/%s/invoice", orderId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	invoice := invoice{}
	json.Unmarshal([]byte(string(body)), &invoice)
	return &invoice, nil
}

func postOrder(c *gin.Context) {
	var newOrder order

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, newOrder)
}
