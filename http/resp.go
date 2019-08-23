package http

import (
	"github.com/gin-gonic/gin"

	"github.com/872409/gatom"
)

func GenJSON(code int, msg string, data interface{}) interface{} {

	resp := gatom.JSON{
		"code": code,
		"msg":  msg,
	}

	if data != nil {
		resp["data"] = data
	}

	return resp
}

func GenSuccessJSON(data interface{}, msg ...string) interface{} {
	_msg := ""
	if len(msg) > 0 {
		_msg = msg[0]
	} else {
		_msg = "ok"
	}

	return GenJSON(1, _msg, data)
}

func GenErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	return GenJSON(code, msg, _data)
}

func Success(c *gin.Context, data interface{}, msg ...string) {
	out := GenSuccessJSON(data, msg...)
	JSON(c, 200, out)
}

func Error(c *gin.Context, msg string, data ...interface{}) {
	JSON(c, 200, GenErrorJSON(msg, -1, data...))
}

func ErrorCode(c *gin.Context, msg string, code int, data ...interface{}) {
	JSON(c, 200, GenErrorJSON(msg, code, data...))
}

func JSON(c *gin.Context, code int, out interface{}) {
	c.JSON(code, out)
}
