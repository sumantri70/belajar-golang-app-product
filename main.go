package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sumantri70/belajar-golang-app-product/src/category/repository"
	"github.com/sumantri70/belajar-golang-app-product/src/common/exception"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sumantri70/belajar-golang-app-product/src/category/controller"
	"github.com/sumantri70/belajar-golang-app-product/src/category/service"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
	"github.com/sumantri70/belajar-golang-app-product/src/database"
	"net/http"
)

func main() {
	db := database.NewDB()
	err := db.Ping()
	helper.PanicIfError(err)

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	createCategoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NexCategoryController(validate, createCategoryService)

	router := httprouter.New()
	router.POST("/v1/category", categoryController.CreateCategory)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:7109",
		Handler: router,
	}

	helper.PanicIfError(server.ListenAndServe())
}
