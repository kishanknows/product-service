package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kishanknows/product-service/internal/handlers"
	"github.com/kishanknows/product-service/internal/middleware"
)

func RegisterProductRoutes(r *gin.Engine) {
	handler := handlers.NewProductHandler()

	public := r.Group("/api/v1")
	{
		public.GET("/product", handler.GetAllProducts)
		public.GET("/product/:id", handler.GetProductById)
	}

	protected := r.Group("/api/v1")
	{
		protected.Use(middleware.AuthMiddleware([]byte(os.Getenv("JWT_SECRET"))))
		protected.POST("/product", handler.CreateProduct)
		protected.DELETE("/product/:id", handler.DeleteProductById)
		protected.PUT("/product/:id", handler.ReplaceProductById)
		protected.PATCH("/product/:id", handler.UpdateProductById)
	}
}