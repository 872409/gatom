package crypto

import (
	"fmt"
	"testing"
)

func TestStrMD5(t *testing.T) {
	str := "dddd.cn"
	cryptoStr := MD5(str)
	fmt.Println(cryptoStr)
}
