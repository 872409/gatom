package util

import (
	"errors"
	"fmt"
	"strconv"
)

func NewCodeError(err string, code int) *codeError {
	return &codeError{errors.New(err), code}
}

func NewCodeErrorF(formatErr string, code int) func(a ...interface{}) error {
	return func(a ...interface{}) error {
		return NewCodeError(fmt.Sprintf(formatErr, a...), code)
	}
}

func ConvertCodeError(err error) CodeError {
	codeErr, ok := err.(CodeError)
	if ok {
		return codeErr
	}
	return nil
}

type CodeError interface {
	error
	Code() int
	Msg() string
}

type codeError struct {
	error
	code int
}

func (e *codeError) Msg() string {
	return e.error.Error()
}

func (e *codeError) Code() int {
	return e.code
}

func (e *codeError) Error() string {
	return "error:" + e.error.Error() + " code:" + strconv.Itoa(e.code)
}
