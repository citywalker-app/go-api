package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/go-playground/validator"
)

func ResetPassword(user *userdomain.User) error {
	if err := validator.New().Struct(user); err != nil {
		return userdomain.ErrBadRequest
	}

	user.SetPassword(user.Password)

	if err := repo.ResetPassword(user); err != nil {
		return err
	}

	return nil
}
