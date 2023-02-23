package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"web_Blogs/pkg/entities"
	"web_Blogs/pkg/repository"
)

type UserUsecase interface {
	GetUserById(id string) (*entities.User, error)
	CreateUser(password, username, gender, facebook, avatar string, bdate time.Time) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(password, username, gender, facebook, avatar string, bdate time.Time) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &entities.User{
		Username:  username,
		Password:  string(hashPassword),
		Gender:    gender,
		Birthdate: bdate,
		Facebook:  facebook,
		Avatar:    avatar,
	}
	return u.repo.Create(user)
}

func (u *userUsecase) GetUserById(id string) (*entities.User, error) {
	return u.repo.FindUserByID(id)
}
