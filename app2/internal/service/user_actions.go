package service

import (
	"app2/internal/models"
	"app2/internal/repository"
	"errors"
)

var (
	ErrIsNotValid    = errors.New("is not valid email")
	ErrAlreadyExists = errors.New("email already exists!")
)

type IUserActionsService interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	CheckUserByEmail(email string) error
}

type userActionsService struct {
	UserActionsRepo repository.IUserActionsRepo
}

func NewUserActionsServiceInit(repo repository.IUserActionsRepo) IUserActionsService {
	return &userActionsService{UserActionsRepo: repo}
}

func (u *userActionsService) CreateUser(user *models.User) error {
	hashPass := GenerateHash(user.Password, user.Salt)
	user.Password = hashPass
	err := u.UserActionsRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userActionsService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.UserActionsRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userActionsService) CheckUserByEmail(email string) error {
	isValid := ValidateEmail(email)
	tempUser, err := u.GetUserByEmail(email)
	if isValid && tempUser == nil {
		return nil
	} else {
		if !isValid {
			return ErrIsNotValid
		} else if tempUser != nil {
			return ErrAlreadyExists
		} else {
			return err
		}
	}
}
