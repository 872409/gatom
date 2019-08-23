package http

import (
	"github.com/gin-gonic/gin"

	"github.com/872409/gatom"
)

type GenJSONHandleFun func(code int, msg string, data interface{}) interface{}

var (
	hasGenJSONHandle = false
	genJSONHandle    GenJSONHandleFun

	JSONCodeName = "code"
	JSONMsgName  = "msg"
	JSONDataName = "data"
)

func SetGenJSONHandler(fun GenJSONHandleFun) {
	hasGenJSONHandle = true
	genJSONHandle = fun
}

func GenJSON(code int, msg string, data interface{}) interface{} {

	json := gatom.JSON{
		JSONCodeName: code,
		JSONMsgName:  msg,
	}

	if data != nil {
		json[JSONDataName] = data
	}

	return json
}

func GenSuccessJSON(data interface{}, msg ...string) interface{} {
	_msg := ""
	if len(msg) > 0 {
		_msg = msg[0]
	} else {
		_msg = "ok"
	}

	code := 1

	if hasGenJSONHandle {
		return genJSONHandle(code, _msg, data)
	}

	return GenJSON(code, _msg, data)
}

func GenErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}

	if code > 0 {
		code = -code
	}

	if hasGenJSONHandle {
		return genJSONHandle(code, msg, _data)
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
