package g

func (g *gContent) JsonSuccess(data interface{}, msg ...string) {
	g.JsonSuccessWithCode(data, 1, msg...)
}

func (g *gContent) JsonSuccessWithCode(data interface{}, code int, msg ...string) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.genSuccessJSON(data, code, msg...))
}

func (g *gContent) JsonError(msg string, data ...interface{}) {
	g.JsonErrorWithCode(msg, -1, data...)
}

func (g *gContent) JsonErrorWithCode(msg string, code int, data ...interface{}) {
	g.AbortWithStatusJSON(200, g.jsonGenerator.genErrorJSON(msg, code, data...))
}
