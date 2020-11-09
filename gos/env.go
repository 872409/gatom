package gos

import (
	"os"
	"strconv"
)

type Env string

const EnvProduction Env = "production"
const EnvLocal Env = "local"

const ENV_GAPP_ENV = "GAPP_ENV"

//
func GetGAppEnv(isLocal bool) Env {
	if isLocal {
		return EnvLocal
	}
	return EnvProduction
}

func GetEnv(key string, def ...string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return def[0]
	}
	return env
}

func GetEnvBool(key string, def ...bool) bool {
	env := os.Getenv(key)

	b, e := strconv.ParseBool(env)
	if e != nil {
		return def[0]
	}
	return b
}
