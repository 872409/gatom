package util

import "os"

func FileIsExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
