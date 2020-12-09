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
	"github.com/872409/gatom/gos"
	"github.com/872409/gatom/log"
	"github.com/872409/gatom/util"
)

var flagLocal = flag.Bool("local", false, "local env. default false")
var flagConfLocalPath = flag.String("conf", "./conf.toml", "config file path")
var flagGAppDebug = flag.Bool("debug", false, "App debug mode. default false")
var flagGAppSignal = flag.Bool("signal", true, "App enabled signal. default true")
var flagGAppPID = flag.Bool("pid", true, "App enabled out PID. default true")

// type BootstrapConfig interface {
// }

// type BootHandle func(bootstrap *Container)

type IContainerOption interface {
	GetContainerOption() *ContainerOption
}

func GetGAppDebug() bool {
	return *flagGAppDebug
}

func DefaultAppOption(name, version string) *ContainerOption {
	env := gos.GetGAppEnv(*flagLocal)

	return &ContainerOption{
		Env:          env,
		Name:         name,
		Version:      version,
		Debug:        *flagGAppDebug,
		EnableSignal: *flagGAppSignal,
		EnablePID:    *flagGAppPID,
	}
}

type ContainerOption struct {
	Version string
	Name    string

	Env          gos.Env
	ConfigLocal  bool
	Debug        bool
	EnablePID    bool
	EnableSignal bool
}

func (receiver *ContainerOption) IsLocal() bool {
	return receiver.Env == gos.EnvLocal
}

func (receiver *ContainerOption) IsProduction() bool {
	return receiver.Env == gos.EnvProduction
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
	OnBeforeBoot func()
	OnBoot       func()
	OnDestroy    func(sig os.Signal, exit bool)
	servers      []ContainerServer
}

// var confLocalEnabled = flag.Bool("confLocal", false, "config load local file, default false")

func (b *Container) InitFromFlag(optionType IContainerOption) {

	if !flag.Parsed() {
		flag.Parse()
	}

	if optionType.GetContainerOption().IsLocal() {
		_confPath := *flagConfLocalPath

		if util.FileIsExist(_confPath) {
			configLoader := config.NewLoader()
			loader := configLoader.LoadFile(_confPath).Unmarshal(&optionType)
			log.Infoln("Init config file error ", loader.Error)
		}
	}
	b.Init(optionType)
}

func (b *Container) Init(option IContainerOption) {

	if option.GetContainerOption().Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

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
	log.Infof("%s run version: '%s' \tdebug:%b \tenablePID:%b enableSignal:%b \tPID:%s \n", option.Name, option.Version, option.Debug, option.EnablePID, option.EnableSignal, pid)
	log.Infof("gatom version:%s\n", gatom.Version)

	// if !option.EnablePID {
	// 	return
	// }

	// err := os.MkdirAll("pids", os.ModePerm)

	err := ioutil.WriteFile(option.Name+".pid", []byte(pid), os.ModePerm)
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
