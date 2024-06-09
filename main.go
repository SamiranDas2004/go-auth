package main

import (
	"github.com/SamiranDas2004/go-auth/dbconnect"
	"github.com/SamiranDas2004/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	dbconnect.ConnectMongoDB()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	routes.Setup(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello World")
	})

	app.Listen(":3000")
}
