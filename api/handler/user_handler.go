package handler

import (
	"net/http"
	"web_Blogs/api/presenter"
	"web_Blogs/pkg/entities"
	"web_Blogs/pkg/usecase/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase      user.UserUsecase
	authencation struct {
		JwtSecret     string
		JwtExpiration int64
	}
}

func NewUserHandler(usecase user.UserUsecase, jwtSecret string, jwtExpiration int64) *UserHandler {
	jwt := new(UserHandler)
	jwt.authencation.JwtSecret = jwtSecret
	jwt.authencation.JwtExpiration = jwtExpiration
	jwt.usecase = usecase
	return jwt
}

// @CreateUser godoc
// @Summary Create new User
// @Description Use for registry
// @Tags Users
// @Accept json
// @Success 200
// @Failure 404
// @Failure 500
// @Router /register [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(entities.UserRequest)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if _, err := handler.usecase.GetUserByUsername(req.Username); err == nil {
		return fiber.NewError(fiber.StatusConflict, "username already existed")
	}
	err := handler.usecase.CreateUser(req.Password, req.Username, req.Gender, req.Facebook, req.Avatar, req.Birthdate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create New user")
	}
	return nil
}

func (handler UserHandler) Authentication(c *fiber.Ctx) error {
	req := new(entities.UserLogin)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := handler.usecase.GetUserByUsername(req.Username)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Email does not exist")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "incorrect password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["exp"] = handler.authencation.JwtExpiration
	tokenString, err := token.SignedString([]byte(handler.authencation.JwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(presenter.UserResponse{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data: &fiber.Map{
			"token": tokenString,
			"user":  user,
		},
	})
}
