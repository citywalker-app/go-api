package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func ResetPassword(user *userdomain.User) error {
	user.SetPassword(user.Password)

	if err := Repo.ResetPassword(user); err != nil {
		return err
	}

	return nil
}
