package gapp

import (
	"github.com/872409/gatom/log"
)

type IHTTPOption interface {
	IContainerOption
	GetHTTPServerOption() *HTTPServerOption
}

func NewHTTP() *HTTPServer {
	container := NewContainer()
	http := &HTTPServer{Container: container}
	container.AddServer(http)
	return http
}

type HTTPServerHandleFun func(httpSvr *HTTPServer)

type HTTPServer struct {
	Container    *Container
	HTTPServer   *GoHTTPServer
	onInitFun    HTTPServerHandleFun
	onDestroyFun HTTPServerHandleFun
}

func (receiver *HTTPServer) GetServerName() string {
	return "http"
}

func (receiver *HTTPServer) OnInit(onInitFun HTTPServerHandleFun) *HTTPServer {
	receiver.onInitFun = onInitFun
	return receiver
}

func (receiver *HTTPServer) OnDestroy(onDestroyFun HTTPServerHandleFun) *HTTPServer {
	receiver.onDestroyFun = onDestroyFun
	return receiver
}

func (receiver *HTTPServer) ServerBoot(bootstrap *Container) {
	receiver.HTTPServer.Boot()
}

func (receiver *HTTPServer) ServerDestroy(bootstrap *Container) {
	receiver.HTTPServer.Stop()

	if receiver.onDestroyFun != nil {
		receiver.onDestroyFun(receiver)
	}
}
// BootWithOption
func (receiver *HTTPServer) BootWithOption(option IHTTPOption) *HTTPServer {
	log.Infoln("HTTPServer.Boot")
	receiver.Container.InitAndBoot(option)
	return receiver
}

// ServerLoad
func (receiver *HTTPServer) ServerLoad(container *Container) {
	log.Infoln("HTTPServer.ServerLoad", container.option)
	option := container.GetOption().(IHTTPOption)

	receiver.HTTPServer = NewGoHTTPServer(option.GetHTTPServerOption())
	receiver.HTTPServer.OnInit = func(http *GoHTTPServer) {
		receiver.onInitFun(receiver)
	}

	receiver.HTTPServer.Init()

}
