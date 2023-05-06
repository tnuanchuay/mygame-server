package main

type SessionMessage struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	ModelId    string  `json:"modelId"`
}

type MovementMessage struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
}

type WebSocketMessage struct {
	Type    int
	Message []byte
	Error   error
}
