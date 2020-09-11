package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Loader struct {
	Viper *viper.Viper
	Error error
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
	l.Error = l.Viper.ReadInConfig()
	if l.Error != nil {
		return l
	}
	l.Error = l.Viper.Unmarshal(conf)
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
