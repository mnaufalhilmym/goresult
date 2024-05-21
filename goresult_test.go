package goresult_test

import (
	"errors"
	"testing"

	"github.com/mnaufalhilmym/goresult"
)

func TestOk(t *testing.T) {
	t.Log("Running TestOk.")

	data := goresult.NewOk("This is a sample Ok value")

	if data.IsOk() {
		if d := data.Unwrap(); d != "This is a sample Ok value" {
			t.Error("TestOk failed. data.Unwrap() should return 'This is a sample Ok value'")
		}
	} else {
		t.Error("TestOk failed. data.IsOk() should return true")
	}
	if data.IsErr() {
		t.Error("TestOk failed. data.IsErr() should not return true")
	}
}

func TestErr(t *testing.T) {
	t.Log("Running TestErr.")

	data := goresult.NewErr[any](errors.New("This is a sample Err value"))

	if data.IsErr() {
		if d := data.UnwrapErr(); d.Error() != "This is a sample Err value" {
			t.Error("TestErr failed. data.UnwrapErr().Error() should return 'This is a sample Err value'")
		}
	} else {
		t.Error("TestErr failed. data.IsErr() should return true")
	}
	if data.IsOk() {
		t.Error("TestErr failed. data.IsOk() should not return true")
	}
}
