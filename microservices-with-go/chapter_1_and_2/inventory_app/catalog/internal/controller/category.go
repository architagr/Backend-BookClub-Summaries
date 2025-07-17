package controller

import (
	"context"

	"inventory.com/catalog/pkg/model"
)

type ICategoryRepository interface {
	Create(ctx context.Context, data *model.Category) (*model.Category, error)
	Update(ctx context.Context, id model.CategoryID, data *model.Category) error
	Get(ctx context.Context, id model.CategoryID) (*model.Category, error)
	GetAll(ctx context.Context) ([]*model.Category, error)
	Delete(ctx context.Context, id model.CategoryID) (*model.Category, error)
}

type CategoryController struct {
	repo ICategoryRepository
}

func NewCategoryController(repo ICategoryRepository) *CategoryController {
	return &CategoryController{repo: repo}
}

func (c *CategoryController) Create(ctx context.Context, data *model.Category) (*model.Category, error) {
	return c.repo.Create(ctx, data)
}

func (c *CategoryController) Update(ctx context.Context, id model.CategoryID, data *model.Category) error {
	return c.repo.Update(ctx, id, data)
}

func (c *CategoryController) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	return c.repo.Get(ctx, id)
}

func (c *CategoryController) GetAll(ctx context.Context) ([]*model.Category, error) {
	return c.repo.GetAll(ctx)
}

func (c *CategoryController) Delete(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	return c.repo.Delete(ctx, id)
}
