package util

import (
	"strconv"
)

type GStr string

func (s GStr) To(def interface{}) (val interface{}) {
	return StrTo(string(s), def)
}

func StrTo(value string, def interface{}) (val interface{}) {

	if value == "" {
		val = def
		return
	}

	var (
		_val interface{}
		err  error
	)

	switch def.(type) {
	case int:
		_val, err = strconv.Atoi(value)
		break
	case bool:
		_val, err = strconv.ParseBool(value)
		break
	case float32:
		_val, err = strconv.ParseFloat(value, 32)
		break
	case float64:
		_val, err = strconv.ParseFloat(value, 64)
		break
	default:
		val = nil
		return
	}

	if err != nil {
		val = def
		return
	}

	val = _val

	return
}

func StrToInt(value string, def ...int) int {

	_def := 0
	if len(def) > 0 {
		_def = def[0]
	}

	return StrTo(value, _def).(int)
}

func StrToInt64(value string, def int64) int64 {
	_val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	return _val
}

func StrArrayToInt64(value []string) []int64 {
	int64Array := make([]int64, 0, len(value))
	for index, val := range value {
		_val, _ := strconv.ParseInt(val, 10, 64)
		int64Array[index] = _val
	}
	return int64Array
}

func StrToFloat32(value string, def float32) float32 {
	_val, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return def
	}
	return float32(_val)
}

func StrToFloat64(value string, def float64) float64 {
	_val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return def
	}
	return _val
}

func StrToBool(value string, def ...bool) bool {

	_def := false

	if len(def) > 0 {
		_def = def[0]
	}

	return StrTo(value, _def).(bool)
}

func StrInArray(array []string, value string) (bool, int) {
	for i, a := range array {
		if a == value {
			return true, i
		}
	}
	return false, -1
}
