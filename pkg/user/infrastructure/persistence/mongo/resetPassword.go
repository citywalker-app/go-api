package mongo

import (
	"context"

	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func (mo *Repository) ResetPassword(user *userdomain.User) error {
	filter := map[string]interface{}{"email": user.Email}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"password": user.Password,
		},
	}

	if _, err := mo.Collection.UpdateOne(context.Background(), filter, update); err != nil {
		return ErrUserNotUpdated
	}

	return nil
}
