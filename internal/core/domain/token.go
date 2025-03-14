// internal/core/domain/token.go

package domain

import "time"

// TokenPair содержит access и refresh токены
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenDetails содержит детальную информацию о токене
type TokenDetails struct {
	AccessToken    string    `json:"-"`
	RefreshToken   string    `json:"-"`
	AccessUuid     string    `json:"-"`
	RefreshUuid    string    `json:"-"`
	AccessExpires  time.Time `json:"-"`
	RefreshExpires time.Time `json:"-"`
	UserId         string    `json:"-"`
}

// TokenClaims представляет содержимое JWT токена
type TokenClaims struct {
	UserId string `json:"user_id"`
	Uuid   string `json:"uuid"`
	Role   string `json:"role"`
}
