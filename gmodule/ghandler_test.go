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

type TE2 struct {
	Changed struct {
		Balance bool
		Freeze  bool
	}
}

func tt2() int {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("tt2 error %v \n", err)
		}
		fmt.Printf("tt2 defer\n")
	}()
	fmt.Printf("tt2 begin\n")
	panic(errors.New("err 123"))
	fmt.Printf("tt2 end\n")
	return 1
}

func TestGHandler_Handle3(t *testing.T) {
	v := tt2()
	fmt.Printf("TestGHandler_Handle3 end %v \n", v)
}
func TestGHandler_Handle2(t *testing.T) {
	t2 := &TE2{}
	t3 := &TE2{}
	t2.Changed.Balance = true
	fmt.Printf("a %v %v \n", t2.Changed, t3.Changed)
}
func TestGHandler_Handle(t *testing.T) {
	p, e := Run(1,
		func(p1 interface{}, p2 interface{}) (interface{}, error) {
			fmt.Printf("a %v %v \n", p1, p2)
			return (p1.(int)) + 1, nil
		},
		func(p1 interface{}, p2 interface{}) (interface{}, error) {
			fmt.Printf("b %v %v \n", p1, p2)
			return (p2.(int)) + 1, nil
		},
		func(p1 interface{}, p2 interface{}) (interface{}, error) {
			fmt.Printf("c %v %v \n", p1, p2)
			return (p2.(int)) + 1, errors.New("cccccc")
		},
		func(p1 interface{}, p2 interface{}) (interface{}, error) {
			fmt.Printf("d %v %v \n", p1, p2)
			return (p2.(int)) + 1, nil
		},
	)
	fmt.Printf("done %v %v \n", p, e)
}
