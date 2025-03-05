package handlers

import "github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"

type UserHandler struct {
	userService ports.UserService
	logger      ports.Logger
}

func NewUserHandler(userService ports.UserService, logger ports.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}
