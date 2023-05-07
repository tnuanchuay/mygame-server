package pubsub

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
)

type PubSub struct {
	messageQueue     chan Message
	subscriberLocker sync.Mutex
	subscribers      map[string][]chan Message
}

var instance *PubSub

func Instance() *PubSub {
	return instance
}

func Init(bufferLength int) {
	if instance != nil {
		return
	}

	instance = New(bufferLength)

	instance.Run()
}

func New(bufferLength int) *PubSub {
	return &PubSub{
		messageQueue:     make(chan Message, bufferLength),
		subscriberLocker: sync.Mutex{},
		subscribers:      make(map[string][]chan Message),
	}
}

func (ps *PubSub) Run() {
	go func() {
		for {
			msg, open := <-ps.messageQueue
			if !open {
				return
			}

			if _, ok := ps.subscribers[msg.Topic]; !ok {
				continue
			}

			for _, listener := range ps.subscribers[msg.Topic] {
				listener <- msg
			}
		}
	}()
}

func (ps *PubSub) GetListenerCount() int {
	ps.subscriberLocker.Lock()
	defer ps.subscriberLocker.Unlock()

	sum := 0
	for _, v := range ps.subscribers {
		sum = sum + len(v)
	}

	return sum
}

func (ps *PubSub) GetAllSubscribersLog() string {
	ps.subscriberLocker.Lock()
	defer ps.subscriberLocker.Unlock()

	sum := "This is list of listeners"
	for topic, subs := range ps.subscribers {
		sum = fmt.Sprintf("%s\nsize: %d\t%s", sum, len(subs), topic)
	}

	return sum
}

func (ps *PubSub) Publish(topic string, data []byte) {
	msg := Message{
		Topic:   topic,
		Payload: data,
	}

	ps.messageQueue <- msg
}

func (ps *PubSub) PublishInterface(topic string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := Message{
		Topic:   topic,
		Payload: b,
	}

	ps.messageQueue <- msg
	return nil
}

func (ps *PubSub) Subscribe(topic string) <-chan Message {
	ps.subscriberLocker.Lock()
	defer ps.subscriberLocker.Unlock()

	log.Println("subscribe", topic)
	sub := make(chan Message, 1000)
	if _, ok := ps.subscribers[topic]; !ok {
		ps.subscribers[topic] = []chan Message{}
	}

	ps.subscribers[topic] = append(ps.subscribers[topic], sub)

	return sub
}

func (ps *PubSub) Unsubscribe(ch <-chan Message) {
	ps.subscriberLocker.Lock()
	defer ps.subscriberLocker.Unlock()

	var keys []string
	for key := range ps.subscribers {
		keys = append(keys, key)
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		for j := 0; j < len(ps.subscribers[key]); j++ {
			sub := ps.subscribers[key][j]
			if sub == ch {
				if len(ps.subscribers[key]) == 1 {
					delete(ps.subscribers, key)
				} else {
					ps.subscribers[key] = append(ps.subscribers[key][:j], ps.subscribers[key][j+1:]...)
				}

				close(sub)
				return
			}
		}
	}
}
