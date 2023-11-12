package service

import (
	"context"
	"github.com/sumantri70/belajar-golang-app-product/src/category/model/request"
	"github.com/sumantri70/belajar-golang-app-product/src/category/model/response"
)

type CreateCategoryService interface {
	Execute(ctx context.Context, request request.CreateCategoryRequest) response.CreateCategoryResponse
	//UpdateCategory(ctx context.Context, categoryId string, request request.UpdateCategoryRequest) response.UpdateCategoryResponse
	//GetCategory(ctx context.Context, categoryId string)
	//DeleteCategory(ctx context.Context, categoryId string)
}
