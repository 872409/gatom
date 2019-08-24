package http

import (
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/872409/gatom"
)

type GenJSONHandleFun func(code int, msg string, data interface{}) gatom.JSON

var (
	hasGenJSONHandle = false
	genJSONHandle    GenJSONHandleFun

	JSONCodeName = "code"
	JSONMsgName  = "msg"
	JSONDataName = "data"
)

func SetGenJSONHandler(handle GenJSONHandleFun) {

	if reflect.TypeOf(handle).Kind() == reflect.Func {
		hasGenJSONHandle = true
		genJSONHandle = handle
	}
}

func GenJSON(code int, msg string, data interface{}) gatom.JSON {

	if hasGenJSONHandle {
		json := genJSONHandle(code, msg, data)

		if json != nil {
			return json
		}
	}

	json := gatom.JSON{
		JSONCodeName: code,
		JSONMsgName:  msg,
	}

	if data != nil {
		json[JSONDataName] = data
	}

	return json
}

func GenSuccessJSON(data interface{}, code int, msg ...string) interface{} {
	var _msg = ""

	if len(msg) > 0 {
		_msg = msg[0]
	} else {
		_msg = "ok"
	}

	if code < 0 {
		code = -code
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

	return GenJSON(code, msg, _data)
}

func Success(c *gin.Context, data interface{}, msg ...string) {
	SuccessCode(c, data, 1, msg...)
}
func SuccessCode(c *gin.Context, data interface{}, code int, msg ...string) {
	out := GenSuccessJSON(data, code, msg...)
	ginJSON(c, 200, out)
}

func Error(c *gin.Context, msg string, data ...interface{}) {
	ginJSON(c, 200, GenErrorJSON(msg, -1, data...))
}

func ErrorCode(c *gin.Context, msg string, code int, data ...interface{}) {
	ginJSON(c, 200, GenErrorJSON(msg, code, data...))
}

func ginJSON(c *gin.Context, code int, out interface{}) {
	c.JSON(code, out)
}
