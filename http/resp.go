package http

import (
	"github.com/gin-gonic/gin"

	"github.com/872409/gatom"
)

func genRespData(code int, msg string, data interface{}) interface{} {

	resp := gatom.JSON{
		"code": code,
		"msg":  msg,
	}

	if data != nil {
		resp["data"] = data
	}

	return resp
}

func GenRespSuccessData(data interface{}, msg ...string) interface{} {
	_msg := ""
	if len(msg) > 0 {
		_msg = msg[0]
	} else {
		_msg = "ok"
	}

	return genRespData(1, _msg, data)
}

func GenRespErrorData(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	return genRespData(code, msg, _data)
}

func RespSuccess(c *gin.Context, data interface{}, msg ...string) {
	out := GenRespSuccessData(data, msg...)
	RespJSON(c, 200, out)
}

func RespError(c *gin.Context, msg string, data ...interface{}) {
	RespJSON(c, 200, GenRespErrorData(msg, -1, data...))
}

func RespErrorCode(c *gin.Context, msg string, code int, data ...interface{}) {
	RespJSON(c, 200, GenRespErrorData(msg, code, data...))
}

func RespJSON(c *gin.Context, code int, out interface{}) {
	c.JSON(code, out)
}
