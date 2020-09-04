package gc

import (
	"github.com/gin-gonic/gin"
)

type GinContent gin.Context

func New(c *gin.Context) *GContext {
	return &GContext{Context: c, jsonGenerator: DefaultGJSONGenerator}
}


type GContext struct {
	*gin.Context
	jsonGenerator *GJSONGenerator
}

func (g *GContext) SetJSONGen(jsonGenerator *GJSONGenerator) *GContext {
	g.jsonGenerator = jsonGenerator
	return g
}

func (g *GinContent) GetP(p string) string {
	return g.Params.ByName(p)
}
