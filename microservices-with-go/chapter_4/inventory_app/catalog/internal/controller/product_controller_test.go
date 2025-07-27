package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"inventory.com/catalog/pkg/model"
)

type MockSubCategoryGetController struct {
	mock.Mock
}

func (m *MockSubCategoryGetController) Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryDetails, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.SubCategoryDetails), args.Error(1)
}

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) Create(ctx context.Context, data *model.ProductBasic) (*model.ProductBasic, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*model.ProductBasic), args.Error(1)
}

func (m *MockProductRepo) Update(ctx context.Context, id model.ProductID, data *model.ProductBasic) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *MockProductRepo) Get(ctx context.Context, id model.ProductID) (*model.ProductBasic, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.ProductBasic), args.Error(1)
}

func (m *MockProductRepo) GetAll(ctx context.Context) ([]*model.ProductBasic, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.ProductBasic), args.Error(1)
}

func (m *MockProductRepo) Delete(ctx context.Context, id model.ProductID) (*model.ProductBasic, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.ProductBasic), args.Error(1)
}

func TestProductController_GetAll(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockSubCategoryCtrl := new(MockSubCategoryGetController)
	ctrl := NewProductController(mockRepo, mockSubCategoryCtrl)

	subCategory := &model.SubCategoryDetails{
		SubCategoryBaseInfo: model.SubCategoryBaseInfo{
			ID:   model.SubCategoryID(1),
			Name: "Personal",
		},
		Category: &model.Category{
			ID:   model.CategoryID(1),
			Name: "Electornics",
		},
	}
	expected := []*model.ProductInformation{
		{ProductBaseInfo: model.ProductBaseInfo{ID: 1, Name: "Laptop"}, SubCategoryDetails: subCategory},
		{ProductBaseInfo: model.ProductBaseInfo{ID: 2, Name: "Phone"}, SubCategoryDetails: subCategory},
	}
	mockRepo.On("GetAll", mock.Anything).Return([]*model.ProductBasic{
		{ProductBaseInfo: model.ProductBaseInfo{ID: 1, Name: "Laptop"}, SubCatID: model.SubCategoryID(1)},
		{ProductBaseInfo: model.ProductBaseInfo{ID: 2, Name: "Phone"}, SubCatID: model.SubCategoryID(1)},
	}, nil)
	mockSubCategoryCtrl.On("Get", mock.Anything, mock.Anything).Return(subCategory, nil)
	result, err := ctrl.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProductController_Delete_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockSubCategoryCtrl := new(MockSubCategoryGetController)
	ctrl := NewProductController(mockRepo, mockSubCategoryCtrl)

	mockRepo.On("Delete", mock.Anything, model.ProductID(404)).Return(&model.ProductBasic{}, errors.New("not found"))

	_, err := ctrl.Delete(context.Background(), 404)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
