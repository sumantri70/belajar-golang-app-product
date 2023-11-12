package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
	web "github.com/sumantri70/belajar-golang-app-product/src/common/model/web/response"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}
	return
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	} else {
		errorsResponse := make([]web.ErrorResponse, 0)
		for _, err := range exception {
			errorsResponse = append(errorsResponse, web.ErrorResponse{
				Field:   err.Field(),
				Message: err.Field() + " " + err.Tag(),
			})
		}
		helper.JsonResponseWrite(writer, helper.ResponseBadRequest(errorsResponse))
		return true
	}
}
