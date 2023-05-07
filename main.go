package main

import (
	"mygame-server/internal/app"
)

func main() {
	app.InitApplicationLog()
	app.InitPubSub()
	app.InitWebServer()
	app.Run()
}
