package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

var playerMovement = make(chan PlayerMovement, 100)

func main() {

	go func() {
		for {
			<-time.After(1 * time.Second)
			log.Println(len(playerMovement))
		}
	}()

	app := fiber.New()
	app.Use("/ws", UpgradeWebsocketMiddleware)

	app.Get("/ws/session", websocket.New(PlayerSession))
	app.Get("/ws/players", websocket.New(GetRoomPlayer))
	app.Get("/ws/move", websocket.New(GetPlayerMovementWSHandler(playerMovement)))
	app.Get("/ws/player/:name", websocket.New(GetPlayerDataByPlayerIdWSHandler(playerMovement)))
	app.Get("/healthcheck", HealthCheckHandler)

	app.Listen(":3000")
}
