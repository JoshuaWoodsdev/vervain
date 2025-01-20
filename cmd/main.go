package main

import (
	"net/http"
	"vervain/internal/config"
	"vervain/internal/middleware"
	"vervain/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.LoadConfig()

	r.Use(middleware.LoggerMiddleware())
	r.Use(CORSMiddleware())

	// Load HTML templates from the "templates" folder
	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Static("/static", "./static")

	r.Run(":" + cfg.Port)
}

// CORSMiddleware handles CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
