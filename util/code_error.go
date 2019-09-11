package util

import (
	"errors"
	"strconv"
)

func NewCodeError(err string, code int) *codeError {
	return &codeError{errors.New(err), code}
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
