package gatom

import (
	"strconv"
)

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
	case bool:
		_val, err = strconv.ParseBool(value)
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

func StrToBool(value string, def ...bool) bool {

	_def := false

	if len(def) > 0 {
		_def = def[0]
	}

	return StrTo(value, _def).(bool)
}
