package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	CreateCategory(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
