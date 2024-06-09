package routes

import (
	"github.com/SamiranDas2004/go-auth/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup returns the Fiber app with defined routes
func Setup(app *fiber.App) {
	// Create a new Fiber app
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/logout", controller.Logout)
}
