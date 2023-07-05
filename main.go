package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	env := os.Getenv("env")

	app := fiber.New()

	// Check Service Status Endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	//Get current environment endpoint
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Env is " + env)
	})

	app.Listen(":3000")
}
