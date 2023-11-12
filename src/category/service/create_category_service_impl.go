package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sumantri70/belajar-golang-app-product/src/category/entity"
	"github.com/sumantri70/belajar-golang-app-product/src/category/model/request"
	"github.com/sumantri70/belajar-golang-app-product/src/category/model/response"
	"github.com/sumantri70/belajar-golang-app-product/src/category/repository"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
)

type CreateCategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CreateCategoryService {
	return &CreateCategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CreateCategoryServiceImpl) Execute(ctx context.Context, request request.CreateCategoryRequest) response.CreateCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	category := toCategory(request)
	category = service.CategoryRepository.Save(ctx, tx, category)

	return toResponse(category)
}

func toCategory(request request.CreateCategoryRequest) entity.Category {
	return entity.Category{
		CategoryId:   uuid.NewString(),
		CategoryName: request.CategoryName,
		Description:  request.Description,
	}
}

func toResponse(category entity.Category) response.CreateCategoryResponse {
	return response.CreateCategoryResponse{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
		Description:  category.Description,
	}
}
