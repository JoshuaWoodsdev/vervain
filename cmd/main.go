package main

import (
	"vervain/internal/config"
	"vervain/internal/middleware"
	"vervain/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.LoadConfig()

	r.Use(middleware.LoggerMiddleware())

	routes.RegisterRoutes(r)

	r.Run(":8080")
	r.Run(":" + cfg.Port)
}
