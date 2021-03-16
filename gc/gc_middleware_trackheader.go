package gc

import "github.com/gin-gonic/gin"

func MiddlewareTrackHeader(trackHeaders ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		for _, header := range trackHeaders {
			context.Header(header, context.GetHeader(header))
		}
	}
}
