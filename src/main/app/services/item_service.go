package services

import (
	"context"
	"ikp_items-api/src/main/app/infrastructure/database"
	"ikp_items-api/src/main/app/model"
)

type IItemService interface {
	Create(createItemRequest *model.CreateItemRequest) (int64, error)
}

type ItemService struct {
	dbClient database.IDbClient
}

func NewItemService(dbClient database.IDbClient) *ItemService {
	return &ItemService{
		dbClient: dbClient,
	}
}

func (i ItemService) Create(createItemRequest *model.CreateItemRequest) (int64, error) {
	result, err := i.dbClient.Context().Item.
		Create().
		SetTitle(createItemRequest.Title).
		Save(context.Background())

	if err != nil {
		return 0, err
	}

	return result.ID, err
}
