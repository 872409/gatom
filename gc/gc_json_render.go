package gc

import "github.com/872409/gatom/util"

func (g *GContext) JSONSuccess(data interface{}, msg ...string) {
	g.JSONSuccessWithCode(data, 1, msg...)
}

func (g *GContext) JSONSuccessWithCode(data interface{}, code int, msg ...string) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenSuccessJSON(data, code, msg...))
}

func (g *GContext) JSONError(msg string, data ...interface{}) {
	g.JSONErrorWithCodeMsg(-1, msg, data...)
}

func (g *GContext) JSONCodeError(err error, data ...interface{}) {
	codeErr, ok := err.(util.CodeError)
	if ok {
		g.JSONErrorWithCodeMsg(codeErr.Code(), codeErr.Msg(), data...)
	} else {
		g.JSONError(err.Error())
	}

}

//
// func (g *GContext) JSONCodeError(err util.CodeError, data ...interface{}) {
// 	g.JSONErrorWithCodeMsg(err.Code(), err.Msg(), data...)
// }

func (g *GContext) JSONErrorWithCode(code int, data ...interface{}) {
	g.JSONErrorWithCodeMsg(code, "error", data...)
}

func (g *GContext) JSONErrorWithCodeMsg(code int, msg string, data ...interface{}) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenErrorJSON(msg, code, data...))
}
