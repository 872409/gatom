package gapp

import (
	"github.com/872409/gatom/log"
)

type IHttpApplicationConfig interface {
	IContainerOption
	GetHttpConfig() HTTPServerOption
}

func NewHttp() *HttpContainer {
	application := NewContainer()
	httpApp := &HttpContainer{Container: application}
	application.AddServer(httpApp)
	return httpApp
}

type HttpContainer struct {
	Container  *Container
	GoHttp     *GoHTTPServer
	OnHttpInit func(http *GoHTTPServer)
}

func (receiver *HttpContainer) GetServerName() string {
	return "http"
}

func (receiver *HttpContainer) ServerBoot(bootstrap *Container) {
	receiver.GoHttp.Boot()
}

func (receiver *HttpContainer) ServerDestroy(bootstrap *Container) {
	receiver.GoHttp.Stop()
}

func (receiver *HttpContainer) BootWithOption(option IHttpApplicationConfig) {
	log.Infoln("HttpContainer.Boot")
	receiver.Container.InitAndBoot(option)
}

//
// func (receiver *HttpContainer) BootWithFlagOption(optionType IHttpApplicationConfig) {
// 	log.Infoln("HttpContainer.BootWithFlagOption")
// 	receiver.Container.InitFromFlagAndBoot(optionType)
// }

func (receiver *HttpContainer) ServerInit(bootstrap *Container) {
	log.Infoln("HttpContainer.ServerInit", bootstrap.option)
	httpApplicationConfig := bootstrap.GetOption().(IHttpApplicationConfig)
	goHttp := NewGoHTTPServer(httpApplicationConfig.GetHttpConfig())
	goHttp.OnInit = receiver.OnHttpInit
	receiver.GoHttp = goHttp
	log.Infoln("HttpContainer.ServerInit", &goHttp.OnInit)
}
