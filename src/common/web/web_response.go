package web

import "net/http"

type Response[T any] struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Errors []ErrorResponse `json:"errors"`
	Data   T               `json:"data"`
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ResponseStatus[T any](response T, code int, status string, errors []ErrorResponse) Response[T] {
	return Response[T]{
		Code:   code,
		Status: status,
		Errors: errors,
		Data:   response,
	}
}

func ResponseOk[T any](response T) Response[T] {
	return ResponseStatus(response, 200, "OK", nil)
}

func ResponseError(code int, status string, errors []ErrorResponse) Response[any] {
	return ResponseStatus[any](nil, code, status, errors)
}

func ResponseBadRequest(errors []ErrorResponse) Response[any] {
	return ResponseError[any](http.StatusBadRequest, "BAD_REQUEST", errors)
}
