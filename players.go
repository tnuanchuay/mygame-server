package main

type Player struct {
	PlayerName string  `json:"playerName"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
}

var roomPlayers map[string]Player = make(map[string]Player)

func mapToList(mp map[string]Player) []Player {
	var lp []Player = []Player{}
	for _, v := range mp {
		lp = append(lp, v)
	}

	return lp
}
