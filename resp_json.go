package gatom

// type JsonRespHandle interface {
// 	GenJSON(code int, msg string, data interface{}) gatom.JSON
// }

type JSONResponseHandler interface {
	JSON(code int, obj interface{})
}

func NewJSONResponse() *JSONResponse {

	return &JSONResponse{
		CodeName: "code",
		MsgName:  "msg",
		DataName: "data",
	}
}

type JSONResponse struct {
	// JsonRespHandle JsonRespHandle
	CodeName string
	MsgName  string
	DataName string
}

func (j *JSONResponse) GenJSON(code int, msg string, data interface{}) JSON {
	json := JSON{
		j.CodeName: code,
		j.MsgName:  msg,
	}

	if data != nil {
		json[j.DataName] = data
	}

	return json
}

func (j *JSONResponse) genSuccessJSON(data interface{}, code int, msg ...string) interface{} {
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

func (j *JSONResponse) genErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}

	if len(data) > 0 {
		_data = data[0]
	}

	if code > 0 {
		code = -code
	}

	return j.GenJSON(code, msg, _data)
}

func (j *JSONResponse) Success(c JSONResponseHandler, data interface{}, msg ...string) {
	j.SuccessCode(c, data, 1, msg...)
}

func (j *JSONResponse) SuccessCode(c JSONResponseHandler, data interface{}, code int, msg ...string) {
	ginJSON(c, 200, j.genSuccessJSON(data, code, msg...))
}

func (j *JSONResponse) Error(c JSONResponseHandler, msg string, data ...interface{}) {
	ginJSON(c, 200, j.genErrorJSON(msg, -1, data...))
}

func (j *JSONResponse) ErrorCode(c JSONResponseHandler, msg string, code int, data ...interface{}) {
	ginJSON(c, 200, j.genErrorJSON(msg, code, data...))
}

func ginJSON(c JSONResponseHandler, code int, out interface{}) {
	c.JSON(code, out)
}
