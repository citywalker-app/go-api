package mongo

import (
	"context"

	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func (mo *Repository) Register(user *userdomain.User) error {
	if _, err := mo.Collection.InsertOne(context.Background(), *user); err != nil {
		return ErrUserNotInserted
	}

	return nil
}
