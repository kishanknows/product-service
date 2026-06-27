package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appErrors "github.com/kishanknows/product-service/internal/errors"
	"github.com/kishanknows/product-service/internal/models"
	"github.com/kishanknows/product-service/internal/response"
	"github.com/kishanknows/product-service/internal/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		service: services.NewProductService(),
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		response.Error(ctx, appErrors.ErrInvalidRequestBody)
		return
	}

	sellerID, ok := ctx.Value("user_id").(int)

	if !ok {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	role, ok := ctx.Value("role").(models.UserRole)
	
	if !ok || (role != models.Admin && role != models.Merchant) {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	product.SellerID = sellerID

	apperr := h.service.CreateProduct(ctx, &product)

	if apperr != nil {
		response.Error(ctx, apperr)
		return
	}

	response.Success(ctx, http.StatusCreated, "product added.", nil)
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.service.GetAllProducts(ctx.Request.Context())

	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "products fetched successfully", products)
}

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := h.service.GetProductById(ctx.Request.Context(), id)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "product fetched successfully", product)
}

func (h *ProductHandler) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	sellerID, ok := ctx.Value("user_id").(int)

	if !ok {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	role, ok := ctx.Value("role").(models.UserRole)
	
	if !ok || (role != models.Admin && role != models.Merchant) {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	err := h.service.DeleteProductById(ctx.Request.Context(), id, sellerID)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, "product deleted successfully", nil)
}

func (h *ProductHandler) ReplaceProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product

	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		response.Error(ctx, appErrors.ErrInvalidRequestBody)
		return
	}

	sellerID, ok := ctx.Value("user_id").(int)

	if !ok {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	role, ok := ctx.Value("role").(models.UserRole)
	
	if !ok || (role != models.Admin && role != models.Merchant) {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	apperr := h.service.ReplaceProductById(ctx.Request.Context(), id, &product, sellerID)

	if apperr != nil {
		response.Error(ctx, apperr)
		return
	}

	response.Success(ctx, http.StatusOK, "product updated.", nil)
}

func (h *ProductHandler) UpdateProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	var req models.UpdateProductRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.Error(ctx, appErrors.ErrInvalidRequestBody)
		return
	}

	sellerID, ok := ctx.Value("user_id").(int)

	if !ok {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	role, ok := ctx.Value("role").(models.UserRole)
	
	if !ok || (role != models.Admin && role != models.Merchant) {
		response.Error(ctx, appErrors.ErrUnauthorized)
		return
	}

	product, apperr := h.service.UpdateProductById(ctx.Request.Context(), id, &req, sellerID)

	if apperr != nil {
		response.Error(ctx, apperr)
		return
	}

	response.Success(ctx, http.StatusOK, "product updated", product)
}

