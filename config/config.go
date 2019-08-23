package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Loader struct {
	Viper *viper.Viper
}

func NewLoader() *Loader {
	l := &Loader{Viper: viper.New()}
	return l
}

func (l *Loader) LoadFile(path string) *Loader {
	l.Viper.SetConfigFile(path)
	return l
}

func (l *Loader) Unmarshal(conf interface{}) *Loader {
	if err := l.Viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("%s", err)
	}
	if err := l.Viper.Unmarshal(conf); err != nil {
		_ = fmt.Errorf("%s", err)
	}
	return l
}

func (l *Loader) Watch(run func(loader *Loader, in fsnotify.Event)) *Loader {
	l.Viper.WatchConfig()
	watch := func(e fsnotify.Event) {
		run(l, e)
	}
	l.Viper.OnConfigChange(watch)
	return l
}
