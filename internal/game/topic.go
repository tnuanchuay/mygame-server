package game

import "fmt"

var (
	TopicPlayerMove         = func(name string) string { return fmt.Sprintf("player/move/%s", name) }
	TopicPlayerDisconnected = func(name string) string { return fmt.Sprintf("player/disconnected/%s", name) }
)
