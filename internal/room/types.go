package room

type Player struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	ModelId    string  `json:"modelId"`
}

type PlayerMovement struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
}
