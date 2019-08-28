package gatom

// type JsonRespHandle interface {
// 	GenJSON(code int, msg string, data interface{}) gatom.JSON
// }

type JSONRenderHandler interface {
	AbortWithStatusJSON(code int, obj interface{})
}

func NewGJSONRender() *gJSONRender {

	return &gJSONRender{
		CodeName:        "code",
		MsgName:         "msg",
		DataName:        "data",
		ErrorStatusCode: 200,
	}
}

type gJSONRender struct {
	CodeName        string
	MsgName         string
	DataName        string
	ErrorStatusCode int
}

func (j *gJSONRender) GenJSON(code int, msg string, data interface{}) JSON {
	json := JSON{
		j.CodeName: code,
		j.MsgName:  msg,
	}

	if data != nil {
		json[j.DataName] = data
	}

	return json
}

func (j *gJSONRender) genSuccessJSON(data interface{}, code int, msg ...string) interface{} {
	var _msg = "ok"

	if len(msg) > 0 {
		_msg = msg[0]
	}

	return j.GenJSON(code, _msg, data)
}

func (j *gJSONRender) genErrorJSON(msg string, code int, data ...interface{}) interface{} {
	var _data interface{}

	if len(data) > 0 {
		_data = data[0]
	}

	return j.GenJSON(code, msg, _data)
}

func (j *gJSONRender) Success(c JSONRenderHandler, data interface{}, msg ...string) {
	j.SuccessWithCode(c, data, 1, msg...)
}

func (j *gJSONRender) SuccessWithCode(c JSONRenderHandler, data interface{}, code int, msg ...string) {
	render(c, 200, j.genSuccessJSON(data, code, msg...))
}

func (j *gJSONRender) Error(c JSONRenderHandler, msg string, data ...interface{}) {
	j.ErrorWithCode(c, msg, -1, data...)
}

func (j *gJSONRender) ErrorWithCode(c JSONRenderHandler, msg string, code int, data ...interface{}) {
	render(c, 200, j.genErrorJSON(msg, code, data...))
}

func render(c JSONRenderHandler, code int, out interface{}) {
	c.AbortWithStatusJSON(code, out)
}
