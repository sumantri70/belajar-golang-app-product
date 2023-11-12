package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sumantri70/belajar-golang-app-product/src/category/entity"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	query := "INSERT INTO category(category_id, category_name, description) values($1,$2,$3)"
	result, err := tx.ExecContext(ctx, query, category.CategoryId, category.CategoryName, category.Description)
	helper.PanicIfError(err)
	fmt.Print(result.RowsAffected())
	return category
}
