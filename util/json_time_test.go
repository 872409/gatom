package util

import (
	"fmt"
	"testing"
)

func TestJSONTime_FormatDate(t1 *testing.T) {
	var date JSONTime = JSONTime{}
	fmt.Println(date.FormatDate())
}
