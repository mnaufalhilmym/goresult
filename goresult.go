package goresult

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Unwrap() T
	UnwrapErr() error
}

type Ok[T any] struct {
	Value T
}

func (Ok[T]) IsOk() bool {
	return true
}

func (Ok[T]) IsErr() bool {
	return false
}

func (o Ok[T]) Unwrap() T {
	return o.Value
}

func (Ok[T]) UnwrapErr() error {
	panic("called `UnwrapErr()` on an `Ok` value")
}

type Err[T any] struct {
	Err error
}

func (Err[T]) IsOk() bool {
	return false
}

func (Err[T]) IsErr() bool {
	return true
}

func (Err[T]) Unwrap() T {
	panic("called `Unwrap()` on an `Err` value")
}

func (e Err[T]) UnwrapErr() error {
	return e.Err
}

func NewOk[T any](value T) Result[T] {
	return Ok[T]{Value: value}
}

func NewErr[T any](err error) Result[T] {
	return Err[T]{Err: err}
}
