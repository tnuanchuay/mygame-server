package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/players", websocket.New(func(conn *websocket.Conn) {
		for {
			conn.WriteJSON(mapToList(roomPlayers))
			<-time.After(1 * time.Millisecond)
		}
	}))

	app.Get("/ws/session", websocket.New(func(conn *websocket.Conn) {
		playerName := ""
		for {
			var msg SessionMessage
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println(playerName, "has left the room")
				delete(roomPlayers, playerName)
				return
			}

			playerName = msg.PlayerName
			roomPlayers[msg.PlayerName] = Player{
				PlayerName: msg.PlayerName,
				X:          msg.X,
				Y:          msg.Y,
			}

			fmt.Println(msg.PlayerName, "has joined the room")
			conn.WriteJSON("ok")
		}
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		fmt.Println("healthcheck")
		return c.JSON("ok")
	})

	app.Listen(":3000")
}
