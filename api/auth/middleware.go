package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JWTMiddleware struct {
	secret []byte
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{secret: []byte(secret)}
}

func (m *JWTMiddleware) Protected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}
		tokenString := authHeader[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return m.secret, nil
		})
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or missing token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("email", claims["sub"].(string))
			return c.Next()
		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or missing token")
		}
	}
}
