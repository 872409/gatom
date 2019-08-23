package config

type AppConfig struct {
	Debug bool
}

func (a *AppConfig) IsReleaseMode() bool {
	return !a.Debug
}
func (a *AppConfig) IsDebugMode() bool {
	return a.Debug
}

type LogConfig struct {
	FilePath string
}

type HttpConfig struct {
	Port int
}
