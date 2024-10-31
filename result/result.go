package result

import (
	"encoding/json"
	"errors"
	"fmt"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dgsys "github.com/darwinOrg/go-common/sys"
)

type Result[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (r *Result[T]) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func (r *Result[T]) ToError() error {
	if !r.Success {
		return dgerr.NewDgError(r.Code, r.Message)
	}

	return nil
}

var simpleSuccess = &Result[*Void]{
	Success: true,
	Code:    0,
}

func Success[T any](data T) *Result[T] {
	return &Result[T]{
		Success: true,
		Code:    0,
		Data:    data,
	}
}

func SimpleSuccess() *Result[*Void] {
	return simpleSuccess
}

func Fail[T any](code int, message string) *Result[T] {
	return &Result[T]{
		Success: false,
		Code:    code,
		Message: message,
	}
}

func SimpleFail[T any](message string) *Result[T] {
	return Fail[T](-1, message)
}

func SimpleFailByError(err error) *Result[*Void] {
	return FailByError[*Void](err)
}

func FailByError[T any](err error) *Result[T] {
	fmt.Println("fail by err: ", err)
	var dgError *dgerr.DgError
	switch {
	case errors.As(err, &dgError):
		return FailByDgError[T](err.(*dgerr.DgError))
	default:
		if dgsys.IsProd() || dgsys.IsPre() {
			return FailByDgError[T](dgerr.SYSTEM_ERROR)
		} else {
			return SimpleFail[T](err.Error())
		}
	}
}

func FailByDgError[T any](err *dgerr.DgError) *Result[T] {
	return &Result[T]{
		Success: false,
		Code:    err.Code,
		Message: err.Message,
	}
}

func ToError[T any](rt *Result[T]) error {
	if rt == nil {
		return dgerr.SYSTEM_ERROR
	}

	return rt.ToError()
}

func ExtractData[T any](rt *Result[T]) (T, error) {
	if rt == nil {
		return *new(T), dgerr.SYSTEM_ERROR
	}

	return rt.Data, rt.ToError()
}
