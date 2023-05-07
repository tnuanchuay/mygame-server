package pubsub

import (
	"reflect"
	"testing"
)

func TestPubSub_UnsubscribeShouldRemoveCorrect(t *testing.T) {
	p := New(1)
	ch1 := p.Subscribe("/1")
	ch2 := p.Subscribe("/1")

	p.Unsubscribe(ch1)

	if _, ok := <-ch1; ok {
		t.Error("ch1 should be close")
	}

	if len(p.subscribers["/1"]) != 1 {
		t.Error("subscriber /1 should have only 1")
	}

	if p.subscribers["/1"][0] != ch2 {
		t.Error("first subscriber of /1 should be ch2")
	}
}

func TestPubSub_UnsubscribeShouldRemoveEntireKeyIfOnly1Subscriber(t *testing.T) {
	p := New(1)

	ch := p.Subscribe("/1")
	p.Unsubscribe(ch)

	if _, ok := p.subscribers["/1"]; ok {
		t.Error("")
	}
}

func TestPubSub_Publish(t *testing.T) {
	p := New(1)
	p.Run()
	ch := p.Subscribe("/1")
	p.Publish("/1", []byte("test"))

	msg := <-ch

	if msg.Topic != "/1" {
		t.Error("topic should be /1")
	}

	if !reflect.DeepEqual(msg.Payload, []byte("test")) {
		t.Error("payload should be byte of test")
	}
}

func TestPubSub_Subscribe(t *testing.T) {
	p := New(1)
	ch := p.Subscribe("/1")

	if ch == nil {
		t.Error("channel should not be nil")
	}

	if len(p.subscribers["/1"]) != 1 {
		t.Error("subscribers map should have /1")
	}

	if p.subscribers["/1"][0] != ch {
		t.Error("first of /1 in subscriber map should be the same that return when call Subscribe()")
	}
}

func TestPubSub_New(t *testing.T) {
	p := New(1)
	if p == nil {
		t.Error("instance should not be nil")
	}

	if cap(p.messageQueue) != 1 {
		t.Error("messageQueue cap should be 1")
	}

	if !p.subscriberLocker.TryLock() {
		t.Error("subscriber locker should be able to lock")
	}

	if p.subscribers == nil {
		t.Error("subscriber repository should not be nil")
	}
}

func TestPubSub_Init(t *testing.T) {
	Init(1)

	if instance == nil {
		t.Error("instance should not be nil")
	}

	if cap(instance.messageQueue) != 1 {
		t.Error("messageQueue cap should be 1")
	}

	if !instance.subscriberLocker.TryLock() {
		t.Error("subscriber locker should be able to lock")
	}

	if instance.subscribers == nil {
		t.Error("subscriber repository should not be nil")
	}
}
