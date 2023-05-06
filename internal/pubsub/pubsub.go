package pubsub

import (
	"encoding/json"
	"fmt"
)

var messageQueue chan Message
var subscribers map[string][]chan Message

func Init(bufferLength int) {
	if messageQueue != nil {
		return
	}

	messageQueue = make(chan Message, bufferLength)
	subscribers = make(map[string][]chan Message)

	go func() {
		for {
			msg := <-messageQueue

			if _, ok := subscribers[msg.Topic]; !ok {
				continue
			}

			for _, listener := range subscribers[msg.Topic] {
				listener <- msg
			}
		}
	}()
}

func GetListenerCount() int {
	sum := 0
	for _, v := range subscribers {
		sum = sum + len(v)
	}

	return sum
}

func GetAllSubscribers() string {
	sum := ""
	for topic, subs := range subscribers {
		for range subs {
			sum = fmt.Sprintf("%s\n%s", sum, topic)
		}
	}

	return sum
}

func Publish(topic string, data []byte) {
	msg := Message{
		Topic:   topic,
		Payload: data,
	}

	messageQueue <- msg
}

func PublishInterface(topic string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := Message{
		Topic:   topic,
		Payload: b,
	}

	messageQueue <- msg
	return nil
}

func Subscribe(topic string) <-chan Message {
	sub := make(chan Message, 1000)
	if _, ok := subscribers[topic]; !ok {
		subscribers[topic] = []chan Message{}
	}

	subscribers[topic] = append(subscribers[topic], sub)

	return sub
}

func Unsubscribe(chan Message) {

}
