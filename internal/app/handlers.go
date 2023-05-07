package app

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	log.Println("healthcheck")
	return c.JSON("ok")
}
