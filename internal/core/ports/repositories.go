package ports

import (
	"context"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) UserService
}
