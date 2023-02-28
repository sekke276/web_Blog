package routes

import (
	"web_Blogs/api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler) {
	users := app.Group("/users")

	users.Post("/register", handler.CreateUser)
}
