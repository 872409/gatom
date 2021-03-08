package gredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisOption struct {
	Host            string
	Pwd             string
	DB              int
	PoolMaxIdle     int
	PoolMaxActive   int
	PoolIdleTimeout int
	ConnectTimeout  int
	ReadTimeout     int
	WriteTimeout    int
}

func DefaultRedisOption() *RedisOption {
	return &RedisOption{
		PoolMaxIdle:     3,
		PoolMaxActive:   500,
		PoolIdleTimeout: 1000,
		ConnectTimeout:  1000,
		ReadTimeout:     1000,
		WriteTimeout:    1000,
	}
}

func (receiver *RedisOption) ApplyDefault() {
	defaultRedisOption := DefaultRedisOption()

	if receiver.PoolMaxIdle <= 0 {
		receiver.PoolMaxIdle = defaultRedisOption.PoolMaxIdle
	}

	if receiver.PoolMaxActive <= 0 {
		receiver.PoolMaxActive = defaultRedisOption.PoolMaxActive
	}

	if receiver.PoolIdleTimeout <= 0 {
		receiver.PoolIdleTimeout = defaultRedisOption.PoolIdleTimeout
	}

	if receiver.ConnectTimeout <= 0 {
		receiver.ConnectTimeout = defaultRedisOption.ConnectTimeout
	}

	if receiver.ReadTimeout <= 0 {
		receiver.ReadTimeout = defaultRedisOption.ReadTimeout
	}

	if receiver.WriteTimeout <= 0 {
		receiver.WriteTimeout = defaultRedisOption.WriteTimeout
	}

}

func NewRedisPool(option *RedisOption) *redis.Pool {

	option.ApplyDefault()

	return &redis.Pool{
		MaxIdle:     option.PoolMaxIdle, // 空闲数
		IdleTimeout: time.Duration(option.PoolIdleTimeout) * time.Millisecond,
		MaxActive:   option.PoolMaxActive, // 最大数
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", option.Host,
				redis.DialPassword(option.Pwd),
				redis.DialDatabase(option.DB),
				redis.DialConnectTimeout(time.Duration(option.ConnectTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(option.ReadTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(option.WriteTimeout)*time.Millisecond))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
