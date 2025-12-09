package server

import (
	"github.com/AliasgharHeidari/test-gorm/internal/api/handler"

	"github.com/gofiber/fiber/v2"
)

func Start() {

	app := fiber.New()

	app.Post("/user", handler.CreateUser)

	app.Listen(":9898")

}
