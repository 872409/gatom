package gatom

import "github.com/gin-gonic/gin"

func NewGContent(c *gin.Context) *gContent {
	return &gContent{Context: c}
}

type gContent struct {
	*gin.Context
}

func (g *gContent) PostInt(name string, def ...int) (val int) {
	_val, _ := g.GetPostForm(name)
	val = StrToInt(_val, def...)
	return
}

func (g *gContent) ParamInt(name string, def ...int) (val int) {
	_val := g.Params.ByName(name)
	val = StrToInt(_val, def...)
	return
}

func (g *gContent) QueryInt(name string, def ...int) (val int) {
	_val := g.Query(name)
	val = StrToInt(_val, def...)
	return
}

func (g *gContent) ParamBoolean(name string, def ...bool) (val bool) {
	_val := g.Params.ByName(name)
	val = StrToBool(_val, def...)
	return
}
