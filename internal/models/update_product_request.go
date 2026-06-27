package models

type UpdateProductRequest struct {
	Name *string `json:"name"`
	Price *int `json:"price"`
}