package service

import (
	"auth-service/internal/domain/entity"
	"context"
)

type TokenDetals struct {
	AccessToken string
	RefreshToken string
	AccessUUID string
	RefreshUUID string
	AtExpires int64
	RtExpires int64
}

//AuthService is the interface for working with authentification
type AuthService interface {
	//registerig user with password
	RegisterWithPassword(ctx context.Context, username, email, password string) (*entity.User, error)

	LoginWithPassword(ctx context.Context, usernameOrEmail, password string) (*entity.User, error)

	RegisterWithSMS(ctx context.Context, phone, ststring) (string, error) //return verification id

	VerifySMSCode(ctx context.Context, verificationID, code string) (*entity.User, error)

	LoginWithSMS(ctx context.Context, phone string) (string, error) //returns verification id

	GenerateOTP(ctx context.Context, userID string) (string, error)

	VerifyOTP(ctx context.Context, userID, otp string) (*TokenDetals, error)

	//Auth with messengers
	InitMessengerAuth(ctx context.Context, messengerType) (string, error) //return url for continue auth

	VerifyMesssengerAuth(ctx context.Context, messengerType, callbackCode string) (*TokenDetals, error)

	RefreshToken(ctx context.Context, refreshToken string) (*TokenDetals, error)

	Logout(ctx context.Context, accessToken) error

	InitPasswordReset(ctx context.Context, email) error

	ConfirmPasswordReset(ctx context.Context, token, newPassword string) error

	//Two factor Auth
	EnableTwoFactor(ctx context.Context, userID, method entity.AuthMethod) (string, error) //returns secret or instruction

	DisableTwoFactor(ctx context.Context, userID) error

	ValidateToken(ctx context.Context, accessToken string) (map[string]interface{}, error)
}
