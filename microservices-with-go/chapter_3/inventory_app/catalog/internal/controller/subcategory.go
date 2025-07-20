package controller

import (
	"context"
	"fmt"

	"inventory.com/catalog/pkg/model"
)

type ISubCategoryRepository interface {
	Create(ctx context.Context, data *model.SubCategoryBasic) (*model.SubCategoryBasic, error)
	Update(ctx context.Context, id model.SubCategoryID, data *model.SubCategoryBasic) error
	Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error)
	GetAll(ctx context.Context) ([]*model.SubCategoryBasic, error)
	Delete(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error)
}

type ICategoryGetController interface {
	Get(ctx context.Context, id model.CategoryID) (*model.Category, error)
}

type SubCategoryController struct {
	repo          ISubCategoryRepository
	catController ICategoryGetController
}

func NewSubCategoryController(repo ISubCategoryRepository, catController ICategoryGetController) *SubCategoryController {
	return &SubCategoryController{
		repo:          repo,
		catController: catController,
	}
}

func (s *SubCategoryController) Create(ctx context.Context, data *model.SubCategoryBasic) (*model.SubCategoryBasic, error) {
	return s.repo.Create(ctx, data)
}

func (s *SubCategoryController) Update(ctx context.Context, id model.SubCategoryID, data *model.SubCategoryBasic) error {
	return s.repo.Update(ctx, id, data)
}

func (s *SubCategoryController) Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryDetails, error) {
	sc, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	cat, err := s.catController.Get(ctx, sc.CatID)
	if err != nil {
		return nil, fmt.Errorf("failed to enrich subcategory with category: %w", err)
	}

	return &model.SubCategoryDetails{
		SubCategoryBaseInfo: model.SubCategoryBaseInfo{
			ID:   sc.ID,
			Name: sc.Name,
		},
		Category: cat,
	}, nil
}

func (s *SubCategoryController) GetAll(ctx context.Context) ([]*model.SubCategoryDetails, error) {
	basics, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.SubCategoryDetails
	for _, b := range basics {
		cat, err := s.catController.Get(ctx, b.CatID)
		if err != nil {
			continue // optionally log
		}
		result = append(result, &model.SubCategoryDetails{
			SubCategoryBaseInfo: model.SubCategoryBaseInfo{
				ID:   b.ID,
				Name: b.Name,
			},
			Category: cat,
		})
	}
	return result, nil
}

func (s *SubCategoryController) Delete(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error) {
	return s.repo.Delete(ctx, id)
}
