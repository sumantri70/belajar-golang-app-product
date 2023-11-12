package command

type Command[R any, T any] interface {
	Execute(request R) T
}
