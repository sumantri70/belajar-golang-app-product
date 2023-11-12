package request

type CreateCategoryRequest struct {
	CategoryName string `json:"categoryName" validate:"required"`
	Description  string `json:"description" validate:"required"`
}
