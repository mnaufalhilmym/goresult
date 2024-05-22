package goresult

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Unwrap() T
	UnwrapErr() error
}

type okResult[T any] struct {
	Value T
}

func (okResult[T]) IsOk() bool {
	return true
}

func (okResult[T]) IsErr() bool {
	return false
}

func (o okResult[T]) Unwrap() T {
	return o.Value
}

func (okResult[T]) UnwrapErr() error {
	panic("called `UnwrapErr()` on an `Ok` value")
}

type errResult[T any] struct {
	Err error
}

func (errResult[T]) IsOk() bool {
	return false
}

func (errResult[T]) IsErr() bool {
	return true
}

func (errResult[T]) Unwrap() T {
	panic("called `Unwrap()` on an `Err` value")
}

func (e errResult[T]) UnwrapErr() error {
	return e.Err
}

func Ok[T any](value T) Result[T] {
	return okResult[T]{Value: value}
}

func Err[T any](err error) Result[T] {
	return errResult[T]{Err: err}
}
