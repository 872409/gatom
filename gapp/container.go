package gapp

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/872409/gatom"
	"github.com/872409/gatom/config"
	"github.com/872409/gatom/log"
	"github.com/872409/gatom/util"
)

// type BootstrapConfig interface {
// }

// type BootHandle func(bootstrap *Container)

type IContainerOption interface {
	GetContainerOption() ContainerOption
}

type ContainerOption struct {
	Version      string
	Name         string
	Debug        bool
	EnablePID    bool
	EnableSignal bool
}

func NewContainer() *Container {

	return &Container{
		servers: []ContainerServer{},
	}
}

type Container struct {


	initOnce    sync.Once
	initialized bool

	option       IContainerOption
	OnFlagParse  func()
	OnBeforeBoot func()
	OnBoot       func()
	OnDestroy    func(sig os.Signal, exit bool)
	servers      []ContainerServer
}

func (b *Container) InitFromFlag(optionType IContainerOption) {

	var confPath string
	flag.StringVar(&confPath, "conf", "./conf.toml", "option path")

	if b.OnFlagParse != nil {
		b.OnFlagParse()
	}

	flag.Parse()

	loader := config.NewLoader().LoadFile(confPath).Unmarshal(&optionType)
	log.Infoln("InitFromFlag error ", loader.Error)
	b.Init(optionType)
}

func (b *Container) Init(option IContainerOption) {

	if b.initialized {
		return
	}

	b.initOnce.Do(func() {
		b.option = option
		b.loadServers()
		b.initialized = true
	})
}

func (b *Container) loadServers() {
	for _, server := range b.servers {
		server.ServerLoad(b)
	}
}

func (b *Container) AddServer(server ContainerServer) {
	b.servers = append(b.servers, server)
}

//
// func (b *Container) InitFromFlagAndBoot(optionType IContainerOption) {
// 	b.InitFromFlag(optionType)
// 	b.Boot()
// }
func (b *Container) InitAndBoot(option IContainerOption) {
	b.InitFromFlag(option)
	b.Boot()
}

func (b *Container) Boot() {
	if !b.initialized {
		panic("container uninitialized")
	}

	if b.OnBeforeBoot != nil {
		b.OnBeforeBoot()
	}

	for _, server := range b.servers {
		server.ServerBoot(b)
	}

	// b.afterServersBoot()

	if b.OnBoot != nil {
		b.OnBoot()
	}

	b.pidFileHandle()
	b.signalHandle()
	// b.afterBoot()
}

func (b *Container) GetOption() IContainerOption {
	return b.option
}

func (b *Container) destroy(sig os.Signal, exit bool) {
	for _, server := range b.servers {
		server.ServerDestroy(b)
	}

	if b.OnDestroy != nil {
		b.OnDestroy(sig, exit)
	}
}

func (b *Container) pidFileHandle() {

	option := b.option.GetContainerOption()
	pid := strconv.Itoa(os.Getpid())
	log.Infof("%s run version: '%s',debug:%s PID:%s \n", option.Name, option.Version, strconv.FormatBool(option.Debug), pid)
	log.Infof("gatom version:%s\n", gatom.Version)

	if !option.EnablePID {
		return
	}

	err := os.MkdirAll("pids", os.ModePerm)

	err = ioutil.WriteFile("pids/"+option.Name+".pid", []byte(pid), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func (b *Container) signalHandle() {
	option := b.option.GetContainerOption()
	if !option.EnableSignal {
		b.destroy(nil, true)
		return
	}

	util.HandleSignal(func(sig os.Signal, exit bool) {
		log.Infof("[Received SIG:%v exit:%s ]", sig, strconv.FormatBool(exit))
		b.destroy(sig, exit)
	})
}
