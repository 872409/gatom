package gapp

import (
	"os"
	"testing"

	"github.com/872409/gatom/log"
)

type TestConfigs struct {
	Option ContainerOption
	Http   HTTPServerOption
}

func (t TestConfigs) GetHttpConfig() HTTPServerOption {
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

	httpApp := NewHttp()
	httpApp.OnHttpInit = func(httpSvr *GoHTTPServer) {
		log.Infoln("App.OnHttpInit")
		// rest.LoadRouters(httpSvr.GinEngine)
	}

	httpApp.Container.OnDestroy = func(sig os.Signal, exit bool) {
		log.Infoln("App", httpApp.Container.option)
	}

	httpApp.BootWithOption(appConfig)
}
