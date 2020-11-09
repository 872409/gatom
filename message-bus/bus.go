package messagebus

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/872409/gatom/log"
)

func NewMessage(topic string, payload interface{}) *Message {
	return &Message{Topic: topic, Payload: payload}
}

type Message struct {
	Topic   string
	Payload interface{}
}

type messageCallback func(msg *Message)

// MessageBus implements publish/subscribe messaging paradigm
type MessageBus interface {
	// Publish publishes arguments to the given topic subscribers
	// Publish block only when the buffer of one of the subscribers is full.
	Publish(topic string, arg interface{})

	PublishMessageTopic(topic string, payload interface{})
	PublishMessage(message *Message)
	SubscribeMessage(topic string, fn messageCallback) unsubscribe
	// Close unsubscribe all handlers from given topic
	Close(topic string)
	// Subscribe subscribes to the given topic
	Subscribe(topic string, fn callback)
	// Unsubscribe unsubscribe handler from the given topic
	Unsubscribe(topic string, fn callback) error
}

type callback func(interface{})
type unsubscribe func() error

type handlersMap map[string][]*handler

type handler struct {
	callback callback
	rvalue   reflect.Value
	queue    chan interface{}
}

type messageBus struct {
	handlerQueueSize int
	mtx              sync.RWMutex
	handlers         handlersMap
}

func (b *messageBus) PublishMessageTopic(topic string, payload interface{}) {
	b.PublishMessage(NewMessage(topic, payload))
}

func (b *messageBus) PublishMessage(msg *Message) {
	b.Publish(msg.Topic, msg)
}

func (b *messageBus) Publish(topic string, arg interface{}) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	if hs, ok := b.handlers[topic]; ok {
		for _, h := range hs {
			h.queue <- arg
		}
	}
}

func (b *messageBus) SubscribeMessage(topic string, fn messageCallback) unsubscribe {
	cb := func(i interface{}) {
		msg := i.(*Message)
		fn(msg)
	}
	b.Subscribe(topic, cb)

	return func() error {
		return b.Unsubscribe(topic, cb)
	}
}

func (b *messageBus) Subscribe(topic string, fn callback) {
	h := &handler{
		callback: fn,
		rvalue:   reflect.ValueOf(fn),
		queue:    make(chan interface{}, b.handlerQueueSize),
	}

	invokeHandlerCallback := func(h *handler, arg interface{}) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorln(fmt.Sprintf("MessageBus.Subscribe.invokeHandlerCallback:%s", r))
			}
		}()
		h.callback(arg)
	}

	go func() {
		for arg := range h.queue {
			invokeHandlerCallback(h, arg)
		}
	}()

	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.handlers[topic] = append(b.handlers[topic], h)
}

func (b *messageBus) Unsubscribe(topic string, fn callback) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	rvalue := reflect.ValueOf(fn)

	if _, ok := b.handlers[topic]; ok {
		for i, h := range b.handlers[topic] {
			if h.rvalue == rvalue {
				close(h.queue)

				b.handlers[topic] = append(b.handlers[topic][:i], b.handlers[topic][i+1:]...)
			}
		}

		return nil
	}

	return fmt.Errorf("topic %s doesn't exist", topic)
}

func (b *messageBus) Close(topic string) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if _, ok := b.handlers[topic]; ok {
		for _, h := range b.handlers[topic] {
			close(h.queue)
		}

		delete(b.handlers, topic)

		return
	}
}

// New creates new MessageBus
// handlerQueueSize sets buffered channel length per subscriber
func NewDefault() MessageBus {
	return New(runtime.NumCPU())
}
func New(handlerQueueSize int) MessageBus {
	if handlerQueueSize == 0 {
		panic("handlerQueueSize has to be greater then 0")
	}

	return &messageBus{
		handlerQueueSize: handlerQueueSize,
		handlers:         make(handlersMap),
	}
}
