package util

import (
	"fmt"
	"testing"
)

func TestCopyFields(t *testing.T) {
	type S1 struct {
		Name string
		Age  int
	}

	type S2 struct {
		Name string
		Age  int
	}

	s1 := S1{"hello", 22}
	var s2 S2
	fmt.Println(s1, s2)
	// CopyFields(&s1, s2)
	// fmt.Println(s1, s2)

	CopyStruct(s1, &s2,"Name")
	fmt.Println(s1, s2)
}
