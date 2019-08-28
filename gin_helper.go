package gatom

import (
	"github.com/gin-gonic/gin"
)

func GinPostInt(c *gin.Context, name string, def ...int) (val int) {
	_val, _ := c.GetPostForm(name)
	val = StrToInt(_val, def...)
	return
}

func GinParamInt(c *gin.Context, name string, def ...int) (val int) {
	_val := c.Params.ByName(name)
	val = StrToInt(_val, def...)
	return
}

func GinQueryInt(c *gin.Context, name string, def ...int) (val int) {
	_val := c.Query(name)
	val = StrToInt(_val, def...)
	return
}

func GinParamBoolean(c *gin.Context, name string, def ...bool) (val bool) {
	_val := c.Params.ByName(name)
	val = StrToBool(_val, def...)
	return
}
