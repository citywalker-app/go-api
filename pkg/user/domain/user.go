package userdomain

import (
	"time"

	"github.com/citywalker-app/go-api/utils"
)

type User struct {
	Email     string    `json:"email" validate:"required,email"`
	FullName  string    `json:"fullName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	Travels   []string  `json:"travels,omitempty" validate:"dive,unique"`
}

func NewUser(email, fullName, password string) *User {
	pass, err := utils.GeneratePassword(password)
	if err != nil {
		return nil
	}
	return &User{
		Email:     email,
		FullName:  fullName,
		Password:  pass,
		CreatedAt: time.Now(),
	}
}

func (u *User) InitializeUser(password ...string) {
	if len(password) > 0 {
		u.SetPassword(password[0])
	} else {
		u.SetPassword(u.Password)
	}
	u.CreatedAt = time.Now()
	u.Travels = make([]string, 0)
}

func (u *User) SetPassword(password string) {
	pass, err := utils.GeneratePassword(password)
	if err != nil {
		return
	}

	u.Password = pass
}

func (u *User) SetTravel(travelID string) {
	u.Travels = append(u.Travels, travelID)
}
