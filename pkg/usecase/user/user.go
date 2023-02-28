package user

import (
	"time"
	"web_Blogs/pkg/entities"
	"web_Blogs/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUserById(id string) (*entities.User, error)
	CreateUser(password, username, gender, facebook, avatar string, bdate time.Time) error
	GetUserByUsername(username string) (*entities.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(password, username, gender, facebook, avatar string, bdate time.Time) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user := &entities.UserRequest{
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

func (u *userUsecase) GetUserByUsername(username string) (*entities.User, error) {
	return u.repo.GetUserByUsername(username)
}
