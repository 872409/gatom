package util

import (
	"fmt"
	"testing"
)

func TestAsCodeError(t *testing.T) {
	err := NewCodeError("AAA", 1)
	as := ConvertCodeError(err)
	fmt.Print("as", as)
}
