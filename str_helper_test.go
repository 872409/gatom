package gatom

import (
	"fmt"
	"testing"
)

func TestStrToInt(t *testing.T) {
	val := StrToInt("a", 2)
	fmt.Println(val)
}

func TestStrToVal(t *testing.T) {
	val := StrTo("10a", 1)
	val2 := StrTo("1", true)
	val3 := StrTo("false", true)
	val4 := StrTo("xx", true)

	fmt.Println(val, val2, val3, val4)
}
