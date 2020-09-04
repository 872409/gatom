package messagebus

import (
	"runtime"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	bus := New(runtime.NumCPU())

	// var wg sync.WaitGroup
	// wg.Add(2)
	bus.Publish("topic", true)

	first := false
	second := false

	bus.Subscribe("topic", func(v interface{}) {
		// defer wg.Done()
		first = v.(bool)
		t.Logf("first:%b", v)
	})

	bus.Subscribe("topic", func(v interface{}) {
		// defer wg.Done()
		second = v.(bool)
		t.Logf("second:%b", v)
	})


	time.Sleep(10 * time.Second)
	// wg.Wait()

	if first == false || second == false {
		t.Fail()
	}
}
