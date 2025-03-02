package domain

type UserError string

const (
	ErrUserNotFound         UserError = "User not found"
	ErrUserAlreadyExists    UserError = "User already exists"
	ErrUserAlreadyActivated UserError = "User already activated"
	ErrUserNotActivated     UserError = "User not activated"
	ErrUserAlreadyLogin     UserError = "User already login"
	ErrUserNotLogin         UserError = "User not login"
	ErrInvalidEmail         UserError = "Invalid email"
	ErrDuplicateEmail       UserError = "Duplicate email"
)
