package app

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
	"mygame-server/internal/game"
	"mygame-server/internal/pubsub"
	"mygame-server/internal/room"
	"time"
)

func PlayerSession(conn *websocket.Conn) {
	defer conn.Close()

	playerName := ""
	for {
		var msg game.SessionMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(playerName, "has left the room")
			room.DeletePlayerByName(playerName)
			pubsub.Instance().PublishInterface(game.TopicPlayerDisconnected(playerName), "")
			return
		}

		playerName = msg.PlayerName
		p := room.Player{
			PlayerName: msg.PlayerName,
			X:          msg.X,
			Y:          msg.Y,
			ModelId:    msg.ModelId,
		}

		room.AddPlayer(p)

		log.Println(msg.PlayerName, "has joined the room")
		conn.WriteJSON("ok")
	}
}

func GetRoomPlayer(conn *websocket.Conn) {
	defer conn.Close()

	conn.WriteJSON(room.GetPlayerList())
	for {
		conn.WriteJSON(room.GetPlayerList())
		<-time.After(1000 * time.Millisecond)
	}
}

func PlayerMovementHandler(conn *websocket.Conn) {
	defer conn.Close()

	for {
		var msg game.MovementMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Errorln(err)
			return
		}

		pm := room.PlayerMovement{
			PlayerName: msg.PlayerName,
			X:          msg.X,
			Y:          msg.Y,
		}

		room.UpdatePlayerPosition(pm)

		err = pubsub.Instance().PublishInterface(game.TopicPlayerMove(pm.PlayerName), pm)
		if err != nil {
			log.Errorln(err)
		}
	}

}

func ListenPlayerMovement(conn *websocket.Conn) {
	defer conn.Close()

	name := conn.Params("name")
	if name == "" {
		return
	}

	playerMovementChan := pubsub.Instance().Subscribe(game.TopicPlayerMove(name))
	defer pubsub.Instance().Unsubscribe(playerMovementChan)

	playerDisconnectedChan := pubsub.Instance().Subscribe(game.TopicPlayerDisconnected(name))
	defer pubsub.Instance().Unsubscribe(playerDisconnectedChan)

	for {
		select {
		case msg := <-playerMovementChan:
			var pm room.PlayerMovement
			err := json.Unmarshal(msg.Payload, &pm)
			if err != nil {
				log.Errorln(err)
			}

			err = conn.WriteJSON(pm)
			if err != nil {
				return
			}

		case <-playerDisconnectedChan:
			return
		}
	}
}
