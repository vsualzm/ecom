package routes

import (
	"ecom/controllers"
	"ecom/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// api.GET("/products", controllers.GetProducts)
	}

	return r
}
