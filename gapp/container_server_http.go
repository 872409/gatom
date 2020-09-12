package gapp

import (
	"github.com/872409/gatom/log"
)

type IHTTPOption interface {
	IContainerOption
	GetHTTPServerOption() HTTPServerOption
}

func NewHTTP() *HTTPServer {
	container := NewContainer()
	http := &HTTPServer{Container: container}
	container.AddServer(http)
	return http
}

type HTTPServer struct {
	Container  *Container
	HTTPServer *GoHTTPServer
	OnInit     func(httpSvr *GoHTTPServer)
	OnDestroy  func()
}

func (receiver *HTTPServer) GetServerName() string {
	return "http"
}

func (receiver *HTTPServer) ServerBoot(bootstrap *Container) {
	receiver.HTTPServer.Boot()
}

func (receiver *HTTPServer) ServerDestroy(bootstrap *Container) {
	receiver.HTTPServer.Stop()

	if receiver.OnDestroy != nil {
		receiver.OnDestroy()
	}
}

func (receiver *HTTPServer) BootWithOption(option IHTTPOption) {
	log.Infoln("HTTPServer.Boot")
	receiver.Container.InitAndBoot(option)
}

//
// func (receiver *HTTPServer) BootWithFlagOption(optionType IHttpApplicationConfig) {
// 	log.Infoln("HTTPServer.BootWithFlagOption")
// 	receiver.Container.InitFromFlagAndBoot(optionType)
// }

func (receiver *HTTPServer) ServerLoad(container *Container) {
	log.Infoln("HTTPServer.ServerLoad", container.option)
	option := container.GetOption().(IHTTPOption)
	httpServer := NewGoHTTPServer(option.GetHTTPServerOption())
	httpServer.OnInit = receiver.OnInit
	httpServer.Init()
	receiver.HTTPServer = httpServer
}
