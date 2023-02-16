package routes

import "github.com/gofiber/fiber/v2"

func UserRouter(app fiber.Router) {
	app.Get("/Login")
}
