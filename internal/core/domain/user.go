// internal/core/domain/user.go

package domain

import (
	"time"
)

// User представляет основную доменную модель пользователя
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Хеш пароля, не сериализуется в JSON
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCredentials модель для аутентификации
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegistration модель для регистрации пользователя
type UserRegistration struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
}
