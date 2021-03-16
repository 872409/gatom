package gredis

import (
	"github.com/gomodule/redigo/redis"
)

func NewSetNotExists(redis *redis.Pool) SetNotExistsInterface {
	return &setNotExists{redis: redis}
}

type SetNotExistsInterface interface {
	SetNotExists(key string, value interface{}, expiredSec int) bool
}

type setNotExists struct {
	redis *redis.Pool
}

func (r setNotExists) SetNotExists(key string, value interface{}, expiredSec int) bool {
	conn := r.redis.Get()
	defer conn.Close()

	// 不存在则设置,并返回成功，存在则返回失败
	reply, _ := redis.Bool(conn.Do("SETNX", key, value))
	// log.Infoln("RequestIDExists SETNX", key, err, "reply:", reply)
	redis.String(conn.Do("SETEX", key, expiredSec))


	if reply {
		// 表示之前不存在
		return false
	}

	return true
}
