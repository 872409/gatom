package util

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	Go(func() {
		fmt.Println("asdfasdfasdfasdfasdfasdfa")
		panic("panicpanicpanicpanic")
	})
}

func TestProtect(t *testing.T) {
	Protect(func() {
		panic("afasfas")
	})
}
