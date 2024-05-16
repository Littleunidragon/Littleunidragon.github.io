package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define login route
	r.POST("/login", func(c *gin.Context) {
		// Get username and password from the form data
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Debugging: Print username and password received
		fmt.Printf("Received username: %s, password: %s\n", username, password)

		// Validate username and password (dummy example)
		if username == "example" && password == "example" {
			// Authentication successful
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			// Debugging: Print message indicating redirection
			fmt.Println("Redirecting to dashboard...")
			// Redirect to the dashboard upon successful login
			c.Redirect(http.StatusFound, "/dashboard")
			return // Return to prevent further execution
		} else {
			// Debugging: Print message indicating authentication failure
			fmt.Println("Authentication failed")
			// Authentication failed
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		}
	})

	// Define create article route
	r.POST("/create-article", func(c *gin.Context) {
		// Handle create article logic here
		// Save article to database, etc.

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
	})

	// Run server
	r.Run(":5500")
}
