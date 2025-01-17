package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Basic route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Vervain!",
		})
	})

	// Start the server
	r.Run(":8080") // Listen on http://localhost:8080
}
