package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var playerMovement = make(chan Player, 100)

func main() {
	app := fiber.New()
	app.Use("/ws", UpgradeWebsocketMiddleware)

	app.Get("/ws/session", websocket.New(PlayerSession))
	app.Get("/ws/players", websocket.New(GetRoomPlayer))
	app.Get("/ws/move", websocket.New(GetPlayerMovementWSHandler(playerMovement)))
	app.Get("/ws/player/:name", websocket.New(GetPlayerDataByPlayerIdWSHandler(playerMovement)))
	app.Get("/healthcheck", HealthCheckHandler)

	app.Listen(":3000")
}
