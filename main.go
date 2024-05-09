package main

import (
	"github.com/SARATHKUMAR-T/fiber-api/configs"
	"github.com/SARATHKUMAR-T/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from fiber & mongoDB"})
	})

	app.Listen(":9000")
}
