package http

import (
	"testing"

	"github.com/872409/gatom"
	"github.com/872409/gatom/log"
)

func genHandle(code int, msg string, data interface{}) gatom.JSON {
	code = 100

	return GenJSON(code, msg, data)
}
func TestLog(t *testing.T) {
	SetGenJSONHandler(genHandle)

	JSONCodeName = "err"
	JSONMsgName = "msg"
	JSONDataName = "result"

	log.Infoln(GenErrorJSON("aa", -1, true))
	log.Infoln(GenSuccessJSON(true, 1, "succeed"))
}
