package gc

import (
	"bytes"
	"io/ioutil"

	"github.com/872409/gatom/log"
)

func (g *GContext) GetBodyString() string {
	buf := make([]byte, g.Request.ContentLength)
	num, _ := g.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	log.Infoln("GetBodyString:" + reqBody)
	g.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody))) // Write body back
	return reqBody
}
