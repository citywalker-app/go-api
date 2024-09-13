package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func Get(user *userdomain.User) error {
	if _, err := repo.GetByEmail(user.Email); err != nil {
		return userdomain.ErrUserNotFound
	}

	return nil
}
