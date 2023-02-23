package handler

import (
	"time"
	"web_Blogs/pkg/usecase/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase user.UserUsecase
}

func NewUserHandler(usecase user.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (hanlder *UserHandler) CreateUser(c *fiber.Ctx) error {
	type createUserRequest struct {
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Gender    string    `json: "gender"`
		Birthdate time.Time `json:"bdate"`
		Avatar    string    `json:"avatar"`
		Facebook  string    `json: "facebook"`
	}

	req := new(createUserRequest)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	err := hanlder.usecase.CreateUser(req.Password, req.Username, req.Gender, req.Facebook, req.Avatar, req.Birthdate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create New user")
	}
	return nil
}
