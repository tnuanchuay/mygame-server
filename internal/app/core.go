package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
	"mygame-server/internal/pubsub"
	"time"
)

var fiberApp *fiber.App

func InitApplicationLog() {
	go func() {
		for {
			<-time.After(1 * time.Second)
			//log.Println("players", room.GetAllPlayerName())
			//log.Println("listener count", pubsub.Instance().GetListenerCount())
			log.Println(pubsub.Instance().GetAllSubscribersLog())
		}
	}()
}

func InitPubSub() {
	pubsub.Init(1000)
}

func InitWebServer() {
	fiberApp = fiber.New()

	fiberApp.Use("/ws", UpgradeWebsocketMiddleware)
	fiberApp.Get("/ws/session", websocket.New(PlayerSession))
	fiberApp.Get("/ws/players", websocket.New(GetRoomPlayer))
	fiberApp.Get("/ws/move", websocket.New(PlayerMovementHandler))
	fiberApp.Get("/ws/player/:name", websocket.New(ListenPlayerMovement))

	fiberApp.Get("/healthcheck", HealthCheckHandler)
}

func Run() {
	fiberApp.Listen(":3000")
}
