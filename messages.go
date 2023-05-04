package main

type SessionMessage struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
}
