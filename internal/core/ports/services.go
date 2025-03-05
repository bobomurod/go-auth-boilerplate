package ports

import (
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, email, password string) error
	ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error
}
