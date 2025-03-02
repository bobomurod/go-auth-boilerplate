package ports

import (
	"context"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
