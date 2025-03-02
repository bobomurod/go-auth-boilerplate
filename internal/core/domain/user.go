package domain

type User struct {
	ID       string
	Email    string
	Password string
	IsActive bool
}

func (u *User) Activate() {
	u.IsActive = true
}
