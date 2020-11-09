package lock

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	redigolib "github.com/gomodule/redigo/redis"
)

func NewLock(config RedisLockConfig) *RedLock {
	pool := redigo.NewPool(&redigolib.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigolib.Conn, error) {
			return redigolib.Dial("tcp", config.RedisHost, redigolib.DialPassword(config.RedisPwd), redigolib.DialDatabase(config.RedisDB))
		},
		TestOnBorrow: func(c redigolib.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	})

	mute := redsync.New(pool)
	return &RedLock{syncMutex: mute}
}

type RunFun func() *RunResult
type RunResult struct {
	Data  interface{}
	Error error
}

type RedLock struct {
	syncMutex *redsync.Redsync
}

func (receiver *RedLock) Run(lockName string, fun RunFun) *RunResult {
	lock := receiver.GetLock(lockName)
	defer lock.Unlock()
	if err := lock.Lock(); err != nil {
		return &RunResult{Data: nil, Error: err}
	}
	return fun()
}
func (receiver *RedLock) GetLock(name string) *redsync.Mutex {
	return receiver.syncMutex.NewMutex(name)
}

func (receiver *RedLock) GetLockTry(name string, tries int) *redsync.Mutex {
	return receiver.syncMutex.NewMutex(name, redsync.WithTries(tries))
}

//
// func main() {
// 	server, err := tempredis.Start(tempredis.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer server.Term()
//
// 	pool := redigo.NewPool(&redigolib.Pool{
// 		MaxIdle:     3,
// 		IdleTimeout: 240 * time.Second,
// 		Dial: func() (redigolib.Conn, error) {
// 			return redigolib.Dial("tcp", server.Socket())
// 		},
// 		TestOnBorrow: func(c redigolib.Conn, t time.Time) error {
// 			_, err := c.Do("PING")
// 			return err
// 		},
// 	})
//
// 	rs := redsync.New(pool)
//
// 	mutex := rs.NewMutex("test-redsync", redsync.WithDriftFactor(1), redsync.WithExpiry(1))
//
// 	if err = mutex.Lock(); err != nil {
// 		panic(err)
// 	}
//
// 	if _, err = mutex.Unlock(); err != nil {
// 		panic(err)
// 	}
// }
