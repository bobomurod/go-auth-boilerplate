package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthMethod string

const (
	AuthMethodPassword AuthMethod = "password"
	AuthMethodSMS      AuthMethod = "sms"
	AuthMethodOTP      AuthMethod = "otp"
	AuthMethodTelegram AuthMethod = "telegram"
	AuthMethodWhatsapp AuthMethod = "whatsapp"
	AuthMethodVK       AuthMethod = "vk"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusBlocked  Status = "blocked"
	StatusPending  Status = "pending"
	StatusIncative Status = "inactive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	Username           string             `bson:"username"`
	Email              string             `bson:"email"`
	Phone              string             `bson:"phone,omitempty"`
	Password           string             `bson:"password,omitempty"`
	PasswordHash       string             `bson:"password_hash,ommitempty"`
	AuthMethods        []AuthMethod       `bson:"auth_methods"`
	Roles              []string           `bson:"roles"`
	Status             Status             `bson:"status"`
	IsTwoFactorEnabled bool               `bson:"is_two_factor_enalbled"`
	LastLogin          time.Time          `bson:"last,omitempty"`
	MessangerIDs       map[string]string  `bson:"messanger_ids,omitempty"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
}

// NewUser создает новый экземпляр пользователя
func NewUser(username, email string, roles []string) *User {
	return &User{
		Username:    username,
		Email:       email,
		Roles:       roles,
		AuthMethods: []AuthMethod{},
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (u *User) SetPassword(passwordHash string) {
	u.PasswordHash = passwordHash
	u.AuthMethods = append(u.AuthMethods, AuthMethodPassword)
	u.UpdatedAt = time.Now()
}

func (u *User) EnableTwoFactor() {
	u.IsTwoFactorEnabled = true
	u.UpdatedAt = time.Now()
}

func (u *User) DisableTwoFactor() {
	u.IsTwoFactorEnabled = false
	u.UpdatedAt = time.Now()
}

func (u *User) block() {
	u.Status = StatusBlocked
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.Status = StatusActive
	u.UpdatedAt = time.Now()
}

func (u *User) AddAuthMethod(method AuthMethod) {
	//check if the method is already added
	for _, m := range u.AuthMethods {
		if m == method {
			return
		}
	}
	u.AuthMethods = append(u.AuthMethods, method)
	u.UpdatedAt = time.Now()
}

func (u *User) AddMessengerID(messengerType, messengerID string) {

}

func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) IsActive() bool {
	return u.Status == StatusActive
}
