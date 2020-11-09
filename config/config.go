package config

import (
	"bytes"

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
	return l.unmarshal(conf)
}

func (l *Loader) UnmarshalFromToml(toml string, conf interface{}) *Loader {
	return l.UnmarshalFromString("toml", toml, conf)
}

func (l *Loader) UnmarshalFromJSON(json string, conf interface{}) *Loader {
	return l.UnmarshalFromString("json", json, conf)
}

func (l *Loader) UnmarshalFromString(configType string, value string, conf interface{}) *Loader {
	l.Viper.SetConfigType(configType)
	l.Error = l.Viper.ReadConfig(bytes.NewBuffer([]byte(value)))
	if l.Error != nil {
		return l
	}
	return l.unmarshal(conf)
}

func (l *Loader) unmarshal(conf interface{}) *Loader {
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
