package gapp

import (
	"testing"

	"github.com/872409/gatom/log"
)

type TestConfigs struct {
	Option ContainerOption
	Http   HTTPServerOption
}

func (t TestConfigs) GetHTTPServerOption() HTTPServerOption {
	return t.Http
}
func (t TestConfigs) GetContainerOption() ContainerOption {
	return t.Option
}

func TestApplication_AddServer(t *testing.T) {

	appConfig := &TestConfigs{
		Option: ContainerOption{Name: "Test", Debug: true, EnableSignal: true, EnablePID: false},
		Http:   HTTPServerOption{DebugMode: "debug", Addr: ":7000"},
	}

	httpApp := NewHTTP()
	httpApp.OnInit = func(httpSvr *GoHTTPServer) {
		log.Infoln("App.OnInit")
		// rest.LoadRouters(httpSvr.GinEngine)
	}

	httpApp.OnDestroy = func() {
		log.Infoln("App", httpApp.Container.option)
	}

	httpApp.BootWithOption(appConfig)
}
