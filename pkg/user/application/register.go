package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func Register(user *userdomain.User) error {
	user.InitializeUser()

	if err := Repo.Register(user); err != nil {
		return err
	}

	return nil
}
