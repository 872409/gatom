package util

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestGoRetry(t *testing.T) {
	GoRetry(3, 1000, func() error {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 500)
			if i == 4 {
				return errors.New("aaaa")
			}
		}
		return nil
	})

	time.Sleep(time.Second * 10)
}
