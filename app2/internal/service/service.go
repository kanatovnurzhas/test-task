package service

import "app2/internal/repository"

type Service struct {
	IUserActionsService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		IUserActionsService: NewUserActionsServiceInit(repo),
	}
}
