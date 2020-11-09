package gapp

import (
	"testing"

	"github.com/872409/gatom/log"
)

type TestConfigs struct {
	Option *ContainerOption
	Http   HTTPServerOption
}

func (t TestConfigs) GetHTTPServerOption() HTTPServerOption {
	return t.Http
}
func (t TestConfigs) GetContainerOption() *ContainerOption {
	return t.Option
}

func TestApplication_AddServer(t *testing.T) {

	appConfig := &TestConfigs{
		Option: &ContainerOption{Name: "Test", Debug: true, EnableSignal: true, EnablePID: false},
		Http:   HTTPServerOption{DebugMode: "debug", Addr: ":7000"},
	}

	NewHTTP().
		OnInit(func(httpSvr *HTTPServer) {
			log.Infoln("App.OnInit")
		}).
		OnDestroy(func(httpSvr *HTTPServer) {
			log.Infoln("App", httpSvr.Container.option)
		}).
		BootWithOption(appConfig)
}
