package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Belajar Pemrograman Golang - GoFiber",
		})
	})

	app.Listen(":8000")
}
