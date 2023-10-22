package main

import (
	"fmt"
	"go-mail-api/mail"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	// Import the package using the directory structure
)


type Contact struct {
	Email   string `json:"email"`
	CompanyName string `json:"companyName"`
	FirstName   string `json:"firstName"`

}


func contactRecieved(c *gin.Context){
	var contact Contact

	// Bind the JSON request body to the Contact struct.
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received Contact: CompanyName=%s, FirstName=%s\n", contact.CompanyName, contact.FirstName)
	msg, err := mail.SendMail(mail.Contact(contact))

	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": msg})


}

func main(){
	router := gin.Default()

	// Middleware to check the domain of incoming requests
	router.Use(func(c *gin.Context) {
		// Get the request's host (e.g., "apppillow.com:8080")
		host := c.Request.Host

		// Split the host to get the domain
		parts := strings.Split(host, ":")
		domain := parts[0]

		// Check if the domain is "apppillow.com"
		if domain != "apppillow.com" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Continue handling the request if the domain is correct
		c.Next()
	})

	router.POST("/contact", contactRecieved)
	router.Run(":8080")
}