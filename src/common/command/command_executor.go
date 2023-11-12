package command

import (
	"github.com/go-playground/validator/v10"
	"github.com/sumantri70/belajar-golang-app-product/src/common/helper"
)

func ExecuteCommand[T any, R any, E Command[R, T]](command E, request R, validate *validator.Validate) T {
	err := validate.Struct(request)
	helper.PanicIfError(err)

	return command.Execute(request)
}
