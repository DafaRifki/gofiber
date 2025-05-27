package routers

import (
	"github.com/DafaRifki/gofiber.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteApp(c *fiber.App) {
	c.Get("/", controllers.UserControllerShow)
}
