package controller

import (
	"context"
	"fmt"

	"inventory.com/catalog/pkg/model"
)

type IProductRepository interface {
	Create(ctx context.Context, data *model.ProductBasic) (*model.ProductBasic, error)
	Update(ctx context.Context, id model.ProductID, data *model.ProductBasic) error
	Get(ctx context.Context, id model.ProductID) (*model.ProductBasic, error)
	GetAll(ctx context.Context) ([]*model.ProductBasic, error)
	Delete(ctx context.Context, id model.ProductID) (*model.ProductBasic, error)
}

type ISubCategoryGetController interface {
	Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryDetails, error)
}

type ProductController struct {
	repo                  IProductRepository
	subCategoryController ISubCategoryGetController
}

func NewProductController(repo IProductRepository, subCategoryController ISubCategoryGetController) *ProductController {
	return &ProductController{
		repo:                  repo,
		subCategoryController: subCategoryController,
	}
}

func (p *ProductController) Create(ctx context.Context, data *model.ProductBasic) (*model.ProductBasic, error) {
	return p.repo.Create(ctx, data)
}

func (p *ProductController) Update(ctx context.Context, id model.ProductID, data *model.ProductBasic) error {
	return p.repo.Update(ctx, id, data)
}

func (p *ProductController) Get(ctx context.Context, id model.ProductID) (*model.ProductInformation, error) {
	pb, err := p.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	subCat, err := p.subCategoryController.Get(ctx, pb.SubCatID)
	if err != nil {
		return nil, fmt.Errorf("failed to enrich product with subcategory: %w", err)
	}

	return &model.ProductInformation{
		ProductBaseInfo: model.ProductBaseInfo{
			ID:           pb.ID,
			Name:         pb.Name,
			Description:  pb.Description,
			Manufacturer: pb.Manufacturer,
			ListCost:     pb.ListCost,
		},
		SubCategoryDetails: subCat,
	}, nil
}

func (p *ProductController) GetAll(ctx context.Context) ([]*model.ProductInformation, error) {
	all, err := p.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.ProductInformation
	for _, pb := range all {
		subCat, err := p.subCategoryController.Get(ctx, pb.SubCatID)
		if err != nil {
			continue // optionally log
		}
		result = append(result, &model.ProductInformation{
			ProductBaseInfo: model.ProductBaseInfo{
				ID:           pb.ID,
				Name:         pb.Name,
				Description:  pb.Description,
				Manufacturer: pb.Manufacturer,
				ListCost:     pb.ListCost,
			},
			SubCategoryDetails: subCat,
		})
	}
	return result, nil
}

func (p *ProductController) Delete(ctx context.Context, id model.ProductID) (*model.ProductBasic, error) {
	return p.repo.Delete(ctx, id)
}
