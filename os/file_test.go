package os

import "testing"

func TestIsExist(t *testing.T) {
	if !FileIsExist(".") {
		t.Error(". must exist")
		return
	}
	if FileIsExist("./aa") {
		t.Error("./aa must not exist")
		return
	}
}
