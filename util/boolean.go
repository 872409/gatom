package util

func B2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func I2B(i int) bool {
	if i > 0 {
		return true
	}
	return false
}
