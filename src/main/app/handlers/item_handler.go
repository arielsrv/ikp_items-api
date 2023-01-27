package handlers

import (
	"ikp_items-api/src/main/app/helpers/ensure"
	"ikp_items-api/src/main/app/model"
	"ikp_items-api/src/main/app/server"
	"ikp_items-api/src/main/app/server/errors"
	"ikp_items-api/src/main/app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type IItemHandler interface {
	Create(ctx *fiber.Ctx) error
}

type ItemHandler struct {
	itemService services.IItemService
}

func NewItemHandler(itemService services.IItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (i ItemHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateItemRequest)
	if err := ctx.BodyParser(request); err != nil {
		return errors.NewError(http.StatusBadRequest, "bad request error, missing key and value properties")
	}

	err := ensure.NotEmpty(request.Title, "bad request error, missing title")
	if err != nil {
		return err
	}

	result, err := i.itemService.Create(request)
	if err != nil {
		return err
	}

	return server.SendCreated(ctx, result)
}
