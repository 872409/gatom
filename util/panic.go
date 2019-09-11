package util

import (
	"fmt"
	"runtime"
)

const (
	maxStack  = 20
	separator = "=====================================\n"
)

var panicHandler func(string)

func OnPanic(h func(string)) {
	panicHandler = func(str string) {
		defer func() {
			recover()
		}()
		h(str)
	}
}

func HandlePanicError(err interface{}, beginStack int) {
	errStr := fmt.Sprintf("\n%sruntime error: %v\ntraceback:\n", separator, err)

	for x := beginStack; x < maxStack; x++ {
		pc, file, line, ok := runtime.Caller(x)
		if !ok {
			break
		}
		errStr += fmt.Sprintf("    stack: %d %v [file: %s] [func: %s] [line: %d]\n", x-1, ok, file, runtime.FuncForPC(pc).Name(), line)
	}

	errStr += separator

	if panicHandler != nil {
		panicHandler(errStr)
	} else {
		fmt.Println(errStr)
	}
}

func HandlePanic() {
	if err := recover(); err != nil {
		HandlePanicError(err, 2)
	}
}
