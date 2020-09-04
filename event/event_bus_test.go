package event

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var eb = NewEventBus()

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publishTo(topic string, data interface{}) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func listenUserEvent() {
	ch1 := make(chan DataEvent)
	eb.Subscribe("user:login", ch1)
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		}
	}
}

func TestEventBus_Publish(t *testing.T) {

	// ch1 := make(chan DataEvent)
	// // ch2 := make(chan DataEvent)
	// ch3 := make(chan DataEvent)
	//
	// eb.Subscribe("topic1", ch1)
	// // eb.Subscribe("topic2", ch2)
	// eb.Subscribe("topic2", ch3)
	//
	// go publishTo("topic1", 1)
	go publishTo("user:login", "Welcome to topic 2")

	listenUserEvent()
	// for {
	// 	select {
	// 	case d := <-ch1:
	// 		go printDataEvent("ch1", d)
	// 	case d := <-ch3:
	// 		go printDataEvent("ch3", d)
	// 	}
	// }

}

func TestEventBus_Subscribe(t *testing.T) {

}
