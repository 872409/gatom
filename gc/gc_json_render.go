package gc

func (g *gContent) JSONSuccess(data interface{}, msg ...string) {
	g.JSONSuccessWithCode(data, 1, msg...)
}

func (g *gContent) JSONSuccessWithCode(data interface{}, code int, msg ...string) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenSuccessJSON(data, code, msg...))
}

func (g *gContent) JSONError(msg string, data ...interface{}) {
	g.JSONErrorWithCodeMsg(-1, msg, data...)
}

func (g *gContent) JSONErrorWithCode(code int, data ...interface{}) {
	g.JSONErrorWithCodeMsg(code, "error", data...)
}

func (g *gContent) JSONErrorWithCodeMsg(code int, msg string, data ...interface{}) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenErrorJSON(msg, code, data...))
}
