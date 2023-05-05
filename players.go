package main

import "sync"

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

var roomPlayersLock = sync.Mutex{}
var roomPlayers = make(map[string]Player)

func mapToList(mp map[string]Player) []Player {
	lp := []Player{}
	for _, v := range mp {
		lp = append(lp, v)
	}

	return lp
}

func updatePlayerPosition(player PlayerMovement) {
	roomPlayersLock.Lock()
	rp, ok := roomPlayers[player.PlayerName]
	if !ok {
		return
	}

	rp.X = player.X
	rp.Y = player.Y

	roomPlayers[player.PlayerName] = rp

	roomPlayersLock.Unlock()
}

func getTotalNumberPlayerInTheRoom() int {
	return len(roomPlayers)
}
