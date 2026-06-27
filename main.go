package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kishanknows/product-service/internal/database"
	"github.com/kishanknows/product-service/internal/routes"
)

func main() {
	godotenv.Load(".env")

	r := gin.Default()

	client, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed")
		panic(err)
	}

	defer client.Disconnect(context.Background())

	r.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	routes.RegisterProductRoutes(r)

	r.Run(":8001")
}
