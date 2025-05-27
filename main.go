package main

import (
	"github.com/DafaRifki/gofiber.git/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.RouteApp(app)

	app.Listen(":8000")
}
