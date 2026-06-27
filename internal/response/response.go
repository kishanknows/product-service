package response

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanknows/product-service/internal/errors"
)

type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data any `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, APIResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func Error(ctx *gin.Context, err *errors.AppError) {
	ctx.JSON(err.Code, APIResponse{
		Success: false,
		Error: err.Message,
	})
}