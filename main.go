package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve login.html for GET requests to /login
	e.GET("/login", func(c echo.Context) error {
		return c.File("login.html")
	})

	// Handle login form submissions
	e.POST("/login", func(c echo.Context) error {
		// Retrieve username and password from the form data
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Implement authentication logic here
		// For example, check if the username and password are valid

		// Example authentication logic
		if username == "admin" && password == "password" {
			// Authentication successful
			// Redirect to the main page or any other appropriate location
			return c.Redirect(http.StatusFound, "/index.html")
		} else {
			// Authentication failed
			return c.String(http.StatusUnauthorized, "Invalid username or password")
		}
	})

	// Start the Echo server on port 5500
	e.Logger.Fatal(e.Start("127.0.0.1:5500"))
}
