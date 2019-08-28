package gatom

import (
	"testing"

	"github.com/872409/gatom/log"
)

func TestLog(t *testing.T) {
	jsonResp := NewGJSONRender()
	log.Infoln(jsonResp.genErrorJSON("aa", -1, true))
	log.Infoln(jsonResp.genSuccessJSON(123, 1, "succeed"))
}
