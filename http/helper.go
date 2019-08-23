package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GParamInt(c *gin.Context, name string, def ...int) (val int) {
	_val := c.Params.ByName(name)

	if _val == "" {
		if len(def) > 0 {
			return def[0]
		}

		return 0
	}
	val, _ = strconv.Atoi(_val)
	return
}

func GParamBoolean(c *gin.Context, name string, def ...bool) (val bool) {
	_val := c.Params.ByName(name)

	if _val == "" {
		if len(def) > 0 {
			return def[0]
		}
		return false
	}
	val, _ = strconv.ParseBool(_val)
	return
}
