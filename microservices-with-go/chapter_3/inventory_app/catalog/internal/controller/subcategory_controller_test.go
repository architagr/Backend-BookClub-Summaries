package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"inventory.com/catalog/pkg/model"
)

type MockSubCategoryRepo struct {
	mock.Mock
}

func (m *MockSubCategoryRepo) Create(ctx context.Context, data *model.SubCategoryBasic) (*model.SubCategoryBasic, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*model.SubCategoryBasic), args.Error(1)
}

func (m *MockSubCategoryRepo) Update(ctx context.Context, id model.SubCategoryID, data *model.SubCategoryBasic) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *MockSubCategoryRepo) Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.SubCategoryBasic), args.Error(1)
}

func (m *MockSubCategoryRepo) GetAll(ctx context.Context) ([]*model.SubCategoryBasic, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.SubCategoryBasic), args.Error(1)
}

func (m *MockSubCategoryRepo) Delete(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.SubCategoryBasic), args.Error(1)
}

type MockCategoryGetController struct {
	mock.Mock
}

func (m *MockCategoryGetController) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Category), args.Error(1)
}

func TestSubCategoryController_Create(t *testing.T) {
	mockRepo := new(MockSubCategoryRepo)
	mockCategoryController := new(MockCategoryGetController)
	ctrl := NewSubCategoryController(mockRepo, mockCategoryController)

	expected := &model.SubCategoryBasic{CatID: 1}
	mockRepo.On("Create", mock.Anything, expected).Return(expected, nil)

	result, err := ctrl.Create(context.Background(), expected)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}
