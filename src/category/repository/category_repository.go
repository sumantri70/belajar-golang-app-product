package repository

import (
	"context"
	"database/sql"
	"github.com/sumantri70/belajar-golang-app-product/src/category/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
}
