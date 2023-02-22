package presenter

import (
	"web_Blogs/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func UserSuccessLogin(data *entities.User, cookie fiber.Cookie) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"token":   cookie,
		"message": "Login scucess",
		"data":    data,
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
		"message": "Register success",
		"data":    nil,
	}
}
