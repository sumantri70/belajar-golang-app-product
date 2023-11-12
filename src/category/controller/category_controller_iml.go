package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sumantri70/belajar-golang-app-product/src/category/model/request"
	"github.com/sumantri70/belajar-golang-app-product/src/category/service"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
	"net/http"
)

type CategoryControllerImpl struct {
	Validate              *validator.Validate
	CreateCategoryService service.CreateCategoryService
}

func NexCategoryController(validate *validator.Validate, categoryService service.CreateCategoryService) CategoryController {
	return &CategoryControllerImpl{
		Validate:              validate,
		CreateCategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) CreateCategory(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	createCategoryRequest := request.CreateCategoryRequest{}
	helper.ReadRequestJson(httpRequest, &createCategoryRequest)

	createCategoryResponse := controller.CreateCategoryService.Execute(httpRequest.Context(), createCategoryRequest)
	helper.WriteResponseJson(writer, helper.ResponseOk(createCategoryResponse))
}
