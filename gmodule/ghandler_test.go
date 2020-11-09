package gmodule

import (
	"errors"
	"fmt"
	"testing"
)

type TestDo struct {
}

func (receiver TestDo) do1(a string) error {
	fmt.Println("handle", "do1", a)
	return nil
}

func (receiver TestDo) do2() error {
	fmt.Println("handle", "do2")
	return nil
}

func (receiver TestDo) do3() error {
	fmt.Println("handle", "do3")
	return errors.New("aaa")
}

func (receiver TestDo) handle() {
	err := RunHandles(
		func() error {
			return receiver.do1("aaa")
		},
		receiver.do2)

	if err != nil {
		fmt.Println("handle 1", err)
	}

	err = RunHandles(
		func() error {
			return receiver.do1("bbbbbb")
		},
		receiver.do3,
		receiver.do2)

	if err != nil {
		fmt.Println("handle 2", err)
	}
}

func TestGHandler_Handle(t *testing.T) {
	do := &TestDo{}
	do.handle()
}
