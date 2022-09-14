package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"os"
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

type order struct {
	ID string `json:"id"`
	Date string `json:"date"`
	Total float32 `json:"total"`
}

var orders = []order{
	{ ID: "1", Date: "01-02-2022", Total: 1000.00},
	{ ID: "2", Date: "02-02-2022", Total: 1500.00},
	{ ID: "3", Date: "03-02-2022", Total: 2000.00},
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
			c.IndentedJSON(http.StatusOK, order)
		}
	}
	if(!orderFound) {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func postOrder(c *gin.Context) {
	var newOrder order

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, newOrder)
}
