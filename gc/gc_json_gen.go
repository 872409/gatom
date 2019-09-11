package gc

import (
	"github.com/872409/gatom/util"
)

var DefaultGJSONGenerator = NewGJSONGenerator()

func NewGJSONGenerator() *GJSONGenerator {

	return &GJSONGenerator{
		CodeName:        "code",
		MsgName:         "msg",
		DataName:        "data",
		ErrorStatusCode: 200,
	}
}

type GJSONGenerator struct {
	CodeName        string
	MsgName         string
	DataName        string
	ErrorStatusCode int
}

func (j *GJSONGenerator) GenJSON(code int, msg string, data interface{}) util.JSON {
	json := util.JSON{
		j.CodeName: code,
		j.MsgName:  msg,
	}

	if data != nil {
		json[j.DataName] = data
	}

	return json
}

func (j *GJSONGenerator) GenSuccessJSON(data interface{}, code int, msg ...string) interface{} {
	var _msg = "ok"

	if len(msg) > 0 {
		_msg = msg[0]
	}

	return j.GenJSON(code, _msg, data)
}

func (j *GJSONGenerator) GenErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}

	if len(data) > 0 {
		_data = data[0]
	}

	return j.GenJSON(code, msg, _data)
}
