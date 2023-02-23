package repository

import (
	"web_Blogs/pkg/entities"
)

type UserRepository interface {
	FindUserByID(id string) (*entities.User, error)
	Create(user *entities.User) error
}
