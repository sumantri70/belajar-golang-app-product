package database

import (
	"database/sql"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:pohodeui70@localhost:5432/belajar_golang_app_product?sslmode=disable")
	//db, err := sql.Open("mysql", "root:pohodeui70@tcp(localhost:3307)/belajar_golang_app_product")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
