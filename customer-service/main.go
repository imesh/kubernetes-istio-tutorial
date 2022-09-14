package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/version", getVersion)
	router.GET("/customers", getCustomers)
	router.POST("/customers", postCustomer)
	router.Run("0.0.0.0:8080")
}

func getVersion(c *gin.Context) {
	response := map[string]string{
		"version": os.Getenv("SERVICE_VERSION"),
	}
	c.IndentedJSON(http.StatusOK, response)
}

type customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}
var customers = []customer{
	{ ID: "1", Name: "David Jones", Email: "david.jones@example.com", Address: "1st Street, New York, USA"},
	{ ID: "2", Name: "Peter Neils", Email: "peter.neils@example.com", Address: "2nd Street, Chicago, USA"},
	{ ID: "3", Name: "Newton Norman", Email: "newton.norman@example.com", Address: "3rd Steet, Sydney, Australia"},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func postCustomer(c *gin.Context) {
	var newCustomer customer

	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}
