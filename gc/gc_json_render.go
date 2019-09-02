package gc

func (g *gContent) JSONSuccess(data interface{}, msg ...string) {
	g.JSONSuccessWithCode(data, 1, msg...)
}

func (g *gContent) JSONSuccessWithCode(data interface{}, code int, msg ...string) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenSuccessJSON(data, code, msg...))
}

func (g *gContent) JSONError(msg string, data ...interface{}) {
	g.JSONErrorWithCode(msg, -1, data...)
}

func (g *gContent) JSONErrorWithCode(msg string, code int, data ...interface{}) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.GenErrorJSON(msg, code, data...))
}
