package mongo

import (
	"context"

	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
)

func (mo *Repository) AddTravel(travelID *string, email *string) error {
	filter := map[string]interface{}{"email": *email}

	update := map[string]interface{}{
		"$push": map[string]interface{}{
			"travels": *travelID,
		},
	}

	result, err := mo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return userdomain.ErrUserNotFound
	}

	return nil
}
