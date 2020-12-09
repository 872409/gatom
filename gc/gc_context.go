package gc

import (
	"github.com/gin-gonic/gin"
)

func HandleFunc(fn func(gc *GContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(New(ctx))
	}
}

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

func (g *GContext) GetP(p string) string {
	return g.Params.ByName(p)
}

func (g *GContext) GetUserAgent() string {
	return g.GetHeader("User-Agent")
}
