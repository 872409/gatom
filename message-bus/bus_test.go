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

	first := false
	// second := false

	bus.Subscribe("topic", func(v interface{}) {
		// defer wg.Done()
		first := v.(string)
		t.Logf("first:%v", first)
	})

	// bus.SubscribeMessage("topic2", func(msg *Message) {
	// 	t.Logf("SubscribeMessage:%v", msg)
	// })
	//
	// bus.Subscribe("topic", func(v interface{}) {
	// 	// defer wg.Done()
	// 	second = v.(bool)
	// 	t.Logf("second:%b", v)
	// })
	//
	bus.Publish("topic", true)
	// bus.PublishMessage(&Message{Topic: "topic2", Payload: "vvvvvvv"})
	//
	time.Sleep(10 * time.Second)
	t.Logf("second:%v", first)
	// // wg.Wait()
	//
	// if first == false || second == false {
	// 	t.Fail()
	// }
}
