package routes

import (
	"vervain/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeHandler)
}
