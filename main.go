package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
	"mygame-server/internal/pubsub"
	"mygame-server/internal/room"
	"time"
)

func main() {

	go func() {
		for {
			<-time.After(1 * time.Second)
			log.Println("players", room.GetAllPlayerName())
			log.Println("listener count", pubsub.Instance().GetListenerCount())
			log.Println(pubsub.Instance().GetAllSubscribers())
		}
	}()

	pubsub.Init(1000)

	app := fiber.New()

	app.Use("/ws", UpgradeWebsocketMiddleware)
	app.Get("/ws/session", websocket.New(PlayerSession))
	app.Get("/ws/players", websocket.New(GetRoomPlayer))
	app.Get("/ws/move", websocket.New(PlayerMovementHandler))
	app.Get("/ws/player/:name", websocket.New(ListenPlayerMovement))

	app.Get("/healthcheck", HealthCheckHandler)

	app.Listen(":3000")
}
