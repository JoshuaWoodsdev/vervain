package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Handle the root route ("/") and serve HTML
	r.GET("/", func(c *gin.Context) {
		// Respond with an HTML file or static content
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})
}
