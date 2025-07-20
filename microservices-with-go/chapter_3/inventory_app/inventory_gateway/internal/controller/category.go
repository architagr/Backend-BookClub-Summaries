package controller

import (
	"context"

	"inventory.com/catalog/pkg/model"
)

type ICategoryGateway interface {
	Create(ctx context.Context, data *model.Category) (*model.Category, error)
	Update(ctx context.Context, id model.CategoryID, data *model.Category) (*model.Category, error)
	Get(ctx context.Context, id model.CategoryID) (*model.Category, error)
}
type CategoryController struct {
	gateway ICategoryGateway
}

func NewCategoryController(gateway ICategoryGateway) *CategoryController {
	return &CategoryController{gateway: gateway}
}
func (c *CategoryController) Create(ctx context.Context, data *model.Category) (*model.Category, error) {
	return c.gateway.Create(ctx, data)
}

func (c *CategoryController) Update(ctx context.Context, id model.CategoryID, data *model.Category) (*model.Category, error) {
	return c.gateway.Update(ctx, id, data)
}

func (c *CategoryController) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	return c.gateway.Get(ctx, id)
}
