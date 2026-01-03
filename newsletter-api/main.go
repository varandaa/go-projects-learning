package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

type client struct {
	//letra maiuscula -> modificaveis pelo JSON
	Name  string
	Email string
	role  string
}

var clients []client = make([]client, 0, 10)

func main() {
	router := gin.Default()
	router.POST("/signup", postClientSignup)
	router.GET("/signup", getClientSignup)
	router.Run(":8085")
}

func postClientSignup(c *gin.Context) {
	var newClient client
	newClient.role = "user"
	err := c.BindJSON(&newClient)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m, err := regexp.MatchString(`^[a-zA-Z.0-9]+@[a-zA-Z]+\.[a-zA-Z0-9]{2,3}$`, newClient.Email)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !m {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "E-mail format is wrong"})
		return
	}
	clients = append(clients, newClient)
	log.Println(newClient.Name)
	fmt.Printf("Clients: %v\n", clients)
	c.IndentedJSON(http.StatusCreated, newClient)
}

func getClientSignup(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clients)
}
