package services

import (
	"context"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/domain"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"
)

type userService struct {
	repo   ports.UserRepository
	logger ports.Logger
}

func NewUserService(repo ports.UserRepository, logger ports.Logger) ports.UserService {
	return &userService{
		repo:   repo,
		logger: logger.With("service", "user"),
	}
}

func (us *userService) RegisterUser(ctx context.Context, email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userService) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error {
	//TODO implement me
	panic("implement me")
}
func (us *userService) Create(ctx context.Context, user domain.User) error {
	us.logger.Info("Creating user", "email", user.Email)
	return us.repo.Create(ctx, user)
}
func (us *userService) GetByEmail(ctx context.Context, email string) ports.UserService {
	us.logger.Info("Getting user by email", "email", email)
	return us.repo.GetByEmail(ctx, email)
}
