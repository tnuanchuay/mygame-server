package game

import "fmt"

var (
	TopicPlayerMove = func(name string) string { return fmt.Sprintf("move/player/%s", name) }
)
