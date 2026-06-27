package errors

import "net/http"

var (
	ErrUnauthorized = New(
		http.StatusUnauthorized,
		"user not authenticated",
		nil,
	)
)