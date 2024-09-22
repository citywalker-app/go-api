package mongo

import (
	"context"

	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *Repository) Create(travel *traveldomain.Travel) error {
	result, err := mo.Collection.InsertOne(context.Background(), *travel)
	if err != nil {
		return ErrTravelNotCreated
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return ErrConvertID
	}

	travel.ID = insertedID.Hex()
	travel.Expenses.Items = make([]traveldomain.Expense, 0)

	return nil
}
