package presenter

import (
	"web_Blogs/pkg/entities"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserSuccessLogin(data *entities.User) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"token":   fiber.Cookie,
		"message": string,
		"data":    entities.User,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": err.Error(),
		"data":    nil,
	}
}

func UserRegister(data *entities.User) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"message": string,
		"data":    nil,
	}
}
