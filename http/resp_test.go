package http

import (
	"testing"

	"github.com/872409/gatom/log"
)

func genHandle(code int, msg string, data interface{}) interface{} {
	code = 100

	return GenJSON(code, msg, data)
}
func TestLog(t *testing.T) {
	SetGenJSONHandler(genHandle)

	JSONCodeName = "err"
	JSONMsgName = "msg"
	JSONDataName = "result"

	json := GenErrorJSON("aa", -1, true)
	log.Infoln(json)
}
