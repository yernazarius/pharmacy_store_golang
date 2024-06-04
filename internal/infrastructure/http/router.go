package http

import (
	"pharmacy-store/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(productHandler *handlers.ProductHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/:id", productHandler.GetProductByID)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("/", productHandler.ListProducts)
		}
	}

	return router
}
