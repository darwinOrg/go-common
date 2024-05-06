package result

import (
	"encoding/json"
	"errors"
	"fmt"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dgsys "github.com/darwinOrg/go-common/sys"
)

type ResultML[T any] struct {
	Success     bool   `json:"success"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	MessageCode string `json:"messageCode"`
	Data        T      `json:"data"`
}

func (r *ResultML[T]) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

var simpleSuccessML = &ResultML[*Void]{
	Success: true,
	Code:    0,
}

func SuccessML[T any](data T) *ResultML[T] {
	return &ResultML[T]{
		Success: true,
		Code:    0,
		Data:    data,
	}
}

func SimpleSuccessML() *ResultML[*Void] {
	return simpleSuccessML
}

func FailML[T any](code int, messageCode string) *ResultML[T] {
	return &ResultML[T]{
		Success:     false,
		Code:        code,
		MessageCode: messageCode,
	}
}

func SimpleFailML[T any](messageCode string) *ResultML[T] {
	return FailML[T](-1, messageCode)
}

func FailByErrorML[T any](err error) *ResultML[T] {
	fmt.Println("fail by err: ", err)
	var dgErrorML *dgerr.DgErrorML
	switch {
	case errors.As(err, &dgErrorML):
		return FailByDgErrorML[T](err.(*dgerr.DgErrorML))
	default:
		if dgsys.IsProd() {
			return FailByDgErrorML[T](dgerr.SYSTEM_ERROR_ML)
		} else {
			return SimpleFailML[T](err.Error())
		}
	}
}

func FailByDgErrorML[T any](err *dgerr.DgErrorML) *ResultML[T] {
	return &ResultML[T]{
		Success:     false,
		Code:        err.Code,
		MessageCode: err.MessageCode,
	}
}
