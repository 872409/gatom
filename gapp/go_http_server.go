package gapp

import (
	"context"
	"net/http"
	"time"

	"github.com/fevin/gracehttp"
	"github.com/gin-gonic/gin"

	"github.com/872409/gatom/log"
)

type GoHTTPServerHandleFun func(receiver *GoHTTPServer)

type HTTPServerOption struct {
	DebugMode string
	Addr      string
}

func NewGoHTTPServer(option HTTPServerOption) *GoHTTPServer {
	return &GoHTTPServer{option: option, graceHttp: gracehttp.NewGraceHTTP()}
}

type GoHTTPServer struct {
	graceHttp  *gracehttp.GraceHTTP
	option     HTTPServerOption
	httpServer *http.Server
	GinEngine  *gin.Engine
	OnInit     func(http *GoHTTPServer)
}

func (receiver *GoHTTPServer) Init() {
	debugMode := receiver.option.DebugMode == gin.DebugMode

	log.Infof("HttpServer boot : %s debugMode:%s\r\n", receiver.option.Addr, receiver.option.DebugMode)

	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	receiver.GinEngine = gin.Default()
	if receiver.OnInit == nil {
		panic("receiver OnInit is nil")
	}
	receiver.OnInit(receiver)

}

func (receiver *GoHTTPServer) Mount(fn func(routerGroup gin.RouterGroup)) {
	routGroup:=receiver.GinEngine.RouterGroup
	fn(routGroup)
}

func (receiver *GoHTTPServer) Boot() {
	debugMode := receiver.option.DebugMode == gin.DebugMode

	if debugMode {
		receiver.runDebug()
	} else {
		go receiver.runReload()
	}
}

func (receiver *GoHTTPServer) runDebug() {
	receiver.httpServer = &http.Server{
		Addr:    receiver.option.Addr,
		Handler: receiver.GinEngine,
	}

	go func() {
		if err := receiver.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (receiver *GoHTTPServer) runReload() {

	receiver.httpServer = &http.Server{
		Addr:    receiver.option.Addr,
		Handler: receiver.GinEngine,
	}

	options := &gracehttp.ServerOption{
		HTTPServer: receiver.httpServer,
	}

	receiver.graceHttp.AddServer(options)
	receiver.graceHttp.Run()
}

func (receiver *GoHTTPServer) Stop() {
	if receiver.httpServer == nil {
		return
	}

	defer func() {
		e := recover()
		if e != nil {
			log.Fatal("HttpServer Stop", e)
		}

	}()

	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := receiver.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("HttpServer Shutdown:", err)
	}

	log.Infoln("------HttpServer stop--------", time.Since(now))
}
