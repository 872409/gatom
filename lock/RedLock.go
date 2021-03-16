package lock

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	redigolib "github.com/gomodule/redigo/redis"

	"github.com/872409/gatom/gredis"
)


type LockInterface interface {
	GetLock(name string, opts ...redsync.Option) *redsync.Mutex
	Run(name string, fun RunFun, opts ...redsync.Option) *RunResult
	GetLockTry(name string, tries int) *redsync.Mutex
}

func NewLock(config RedisLockConfig) LockInterface {
	pool := &redigolib.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigolib.Conn, error) {
			return redigolib.Dial("tcp", config.RedisHost, redigolib.DialPassword(config.RedisPwd), redigolib.DialDatabase(config.RedisDB))
		},
		TestOnBorrow: func(c redigolib.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return NewWithPool(pool)
}

func New(option *gredis.RedisOption) *RedLock {
	return NewWithPool(gredis.NewRedisPool(option))
}

func NewWithPool(pool *redigolib.Pool) *RedLock {
	redigoPool := redigo.NewPool(pool)
	mute := redsync.New(redigoPool)
	return &RedLock{syncMutex: mute}
}

// type Option redsync.Option

type RunFun func() *RunResult
type RunResult struct {
	Data  interface{}
	Error error
}

func Result(err error, data ...interface{}) *RunResult {
	result := &RunResult{Error: err}
	if len(data) == 1 {
		result.Data = data[0]
	}
	return result
}

type RedLock struct {
	syncMutex *redsync.Redsync
}

func (receiver *RedLock) Run(lockName string, fun RunFun, opts ...redsync.Option) *RunResult {
	lock := receiver.GetLock(lockName, opts...)
	defer lock.Unlock()
	if err := lock.Lock(); err != nil {
		return &RunResult{Data: nil, Error: err}
	}
	return fun()
}
func (receiver *RedLock) GetLock(name string, opts ...redsync.Option) *redsync.Mutex {
	return receiver.syncMutex.NewMutex(name, opts...)
}

func (receiver *RedLock) GetLockTry(name string, tries int) *redsync.Mutex {
	return receiver.syncMutex.NewMutex(name, redsync.WithTries(tries))
}
