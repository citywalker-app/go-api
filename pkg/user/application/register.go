package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/go-playground/validator"
)

func Register(user *userdomain.User) error {
	if err := validator.New().Struct(user); err != nil {
		return userdomain.ErrBadRequest
	}

	if user.Password == "" {
		return userdomain.ErrBadRequest
	}

	user.InitializeUser()

	if err := repo.Register(user); err != nil {
		return err
	}

	return nil
}
