package room

import (
	"fmt"
	"sync"
)

var roomPlayersLock = sync.Mutex{}
var roomPlayers = make(map[string]Player)

func mapToList(mp map[string]Player) []Player {
	lp := []Player{}
	for _, v := range mp {
		lp = append(lp, v)
	}

	return lp
}

func GetPlayerList() []Player {
	roomPlayersLock.Lock()
	defer roomPlayersLock.Unlock()

	return mapToList(roomPlayers)
}

func UpdatePlayerPosition(player PlayerMovement) {
	roomPlayersLock.Lock()
	defer roomPlayersLock.Unlock()

	rp, ok := roomPlayers[player.PlayerName]
	if !ok {
		return
	}

	rp.X = player.X
	rp.Y = player.Y

	roomPlayers[player.PlayerName] = rp
}

func GetAllPlayerName() string {
	names := ""
	for k, _ := range roomPlayers {
		if names == "" {
			names = fmt.Sprintf("%s", k)
		} else {
			names = fmt.Sprintf("%s, %s", names, k)
		}
	}

	return names
}

func AddPlayer(player Player) {
	roomPlayersLock.Lock()
	defer roomPlayersLock.Unlock()

	roomPlayers[player.PlayerName] = player
}

func DeletePlayerByName(name string) {
	roomPlayersLock.Lock()
	defer roomPlayersLock.Unlock()

	_, ok := roomPlayers[name]
	if !ok {
		return
	}

	delete(roomPlayers, name)
}
