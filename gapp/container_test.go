package gapp

import (
	"os"
	"testing"

	"github.com/872409/gatom/log"
)

type ContainerConfig struct {
	Option *ContainerOption
}

func (t ContainerConfig) GetContainerOption() *ContainerOption {
	return t.Option
}

func TestNewBootstrap(t *testing.T) {

	appConfig := &ContainerConfig{
		Option: &ContainerOption{Name: "Test", Debug: true, EnableSignal: true, EnablePID: false},
	}

	container := NewContainer()

	container.OnBoot = func() {
		log.Infoln("container OnBoot", container.option)
	}

	container.OnDestroy = func(sig os.Signal, exit bool) {
		log.Infoln("container OnDestroy", container.option)
	}

	container.InitAndBoot(appConfig)
}
