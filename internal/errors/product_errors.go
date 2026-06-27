package errors

import "net/http"

var (
	ErrProductNotFound = New(
		http.StatusNotFound,
		"product doesn't exist",
		nil,
	)

	ErrInternalServer = New(
		http.StatusInternalServerError,
		"internal server error",
		nil,
	)

	ErrInvalidProductID = New(
		http.StatusBadRequest,
		"invalid product id",
		nil,
	)

	ErrInvalidRequestBody = New(
		http.StatusBadRequest,
		"invalid request body",
		nil,
	)
)