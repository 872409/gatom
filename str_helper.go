package gatom

import "strconv"

func StrToInt(value string, def ...int) (val int) {

	if value == "" {
		if len(def) > 0 {
			return def[0]
		}

		val = 0
		return
	}

	val, _ = strconv.Atoi(value)
	return
}

func StrToBool(value string, def ...bool) (val bool) {

	if value == "" {
		if len(def) > 0 {
			return def[0]
		}

		val = false
		return
	}

	val, _ = strconv.ParseBool(value)
	return
}
