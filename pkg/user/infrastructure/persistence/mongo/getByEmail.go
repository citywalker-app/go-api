package mongo

import (
	"context"

	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func (mo *Repository) GetByEmail(email string) (*userdomain.User, error) {
	user := userdomain.User{}
	filter := map[string]interface{}{"email": email}

	if err := mo.Collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}
