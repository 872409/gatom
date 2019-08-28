package http

import (
	"github.com/gin-gonic/gin"

	"github.com/872409/gatom"
)

// type JsonRespHandle interface {
// 	GenJSON(code int, msg string, data interface{}) gatom.JSON
// }

func New() *JsonResp {

	return &JsonResp{
		CodeName: "code",
		MsgName:  "msg",
		DataName: "data",
	}
}

type JsonResp struct {
	// JsonRespHandle JsonRespHandle
	CodeName string
	MsgName  string
	DataName string
}

func (j *JsonResp) GenJSON(code int, msg string, data interface{}) gatom.JSON {
	json := gatom.JSON{
		j.CodeName: code,
		j.MsgName:  msg,
	}

	if data != nil {
		json[j.DataName] = data
	}

	return json
}

func (j *JsonResp) genSuccessJSON(data interface{}, code int, msg ...string) interface{} {
	var _msg = ""

	if len(msg) > 0 {
		_msg = msg[0]
	} else {
		_msg = "ok"
	}

	if code < 0 {
		code = -code
	}

	return j.GenJSON(code, _msg, data)
}

func (j *JsonResp) genErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}

	if len(data) > 0 {
		_data = data[0]
	}

	if code > 0 {
		code = -code
	}

	return j.GenJSON(code, msg, _data)
}

func (j *JsonResp) Success(c *gin.Context, data interface{}, msg ...string) {
	j.SuccessCode(c, data, 1, msg...)
}

func (j *JsonResp) SuccessCode(c *gin.Context, data interface{}, code int, msg ...string) {
	ginJSON(c, 200, j.genSuccessJSON(data, code, msg...))
}

func (j *JsonResp) Error(c *gin.Context, msg string, data ...interface{}) {
	ginJSON(c, 200, j.genErrorJSON(msg, -1, data...))
}

func (j *JsonResp) ErrorCode(c *gin.Context, msg string, code int, data ...interface{}) {
	ginJSON(c, 200, j.genErrorJSON(msg, code, data...))
}

func ginJSON(c *gin.Context, code int, out interface{}) {
	c.JSON(code, out)
}
