package services

import "github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"

type UserService struct {
	repo   ports.UserRepository
	logger ports.Logger
}

func NewUserSerive(repo ports.UserRepository, logger ports.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}
