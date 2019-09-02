package gc

import (
	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) *gContent {
	return &gContent{Context: c, jsonGenerator: DefaultGJSONGenerator}
}

type gContent struct {
	*gin.Context
	jsonGenerator *GJSONGenerator
}

func (g *gContent) SetJSONGen(jsonGenerator *GJSONGenerator) {
	g.jsonGenerator = jsonGenerator
}
