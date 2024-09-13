package userapplication

import (
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func Login(user *userdomain.User) error {
	foundedUser, err := repo.GetByEmail(user.Email)
	if err != nil {
		return userdomain.ErrUserNotFound
	}

	if areEquals := bcrypt.CompareHashAndPassword([]byte(foundedUser.Password), []byte(user.Password)); areEquals != nil {
		return userdomain.ErrWrongCredentials
	}

	*user = *foundedUser

	return nil
}
