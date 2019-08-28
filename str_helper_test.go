package gatom

import (
	"fmt"
	"testing"
)

func TestStrToInt(t *testing.T) {
	val := StrToInt("a", 2)
	fmt.Println(val)
}

type A struct {
	Name string
}

func TestStrToVal(t *testing.T) {

	ms := map[string]A{"aa": {Name: "N"}}

	fmt.Println("aa", ms, ms["aaaa"].Name)

	//
	// val := StrTo("10a", 1)
	// val2 := StrTo("1", true)
	// val3 := StrTo("false", true)
	// val4 := StrTo("xx", true)
	//
	// fmt.Println(val, val2, val3, val4)
}
