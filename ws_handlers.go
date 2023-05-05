package main

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

func PlayerSession(conn *websocket.Conn) {
	defer conn.Close()

	playerName := ""
	for {
		var msg SessionMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(playerName, "has left the room")
			delete(roomPlayers, playerName)
			return
		}

		playerName = msg.PlayerName
		roomPlayers[msg.PlayerName] = Player{
			PlayerName: msg.PlayerName,
			X:          msg.X,
			Y:          msg.Y,
		}

		log.Println(msg.PlayerName, "has joined the room")
		conn.WriteJSON("ok")
	}
}

func GetRoomPlayer(conn *websocket.Conn) {
	defer conn.Close()

	conn.WriteJSON(mapToList(roomPlayers))
	for {
		conn.WriteJSON(mapToList(roomPlayers))
		<-time.After(1000 * time.Millisecond)
	}
}

func GetPlayerMovementWSHandler(playerMovement chan<- Player) func(conn *websocket.Conn) {
	return func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			var msg MovementMessage
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println(err)
				return
			}

			p := Player{
				msg.PlayerName,
				msg.X,
				msg.Y,
			}
			playerMovement <- p
			updatePlayerPosition(p)
			log.Println(p.PlayerName, "is moving to", p.X, p.Y)
		}
	}
}

func GetPlayerDataByPlayerIdWSHandler(playerMovement chan Player) func(conn *websocket.Conn) {

	return func(conn *websocket.Conn) {
		defer conn.Close()

		name := conn.Params("name")

		for {
			p := <-playerMovement
			if p.PlayerName != name {
				playerMovement <- p
				continue
			}

			p, ok := roomPlayers[p.PlayerName]
			if !ok {
				return
			}

			conn.WriteJSON(p)
		}
	}
}
