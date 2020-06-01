package gc

import (
	"github.com/gin-gonic/gin"
)

type GinContent gin.Context

func New(c *gin.Context) *gContent {
	return &gContent{Context: c, jsonGenerator: DefaultGJSONGenerator}
}

type gContent struct {
	*gin.Context
	jsonGenerator *GJSONGenerator
}

func (g *gContent) SetJSONGen(jsonGenerator *GJSONGenerator) *gContent {
	g.jsonGenerator = jsonGenerator
	return g
}

func (g *GinContent) GetP(p string) string {
	return g.Params.ByName("")
}
