package game

import "testing"

func TestTopicPlayerMoveShouldReturnCorrectTopic(t *testing.T) {
	topic := TopicPlayerMove("bin")

	if topic != "player/move/bin" {
		t.Error("topic should be player/move/bin")
	}
}

func TestTopicPlayerDisconnectedSholdReturnCorrectTopic(t *testing.T) {
	topic := TopicPlayerDisconnected("bin")

	if topic != "player/disconnected/bin" {
		t.Error("topic should be player/disconnected/bin")
	}
}
