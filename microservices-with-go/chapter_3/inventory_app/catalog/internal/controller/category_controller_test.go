package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"inventory.com/catalog/internal/controller"
	"inventory.com/catalog/pkg/model"
)

// --- Mocks ---
type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) Create(ctx context.Context, data *model.Category) (*model.Category, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) Update(ctx context.Context, id model.CategoryID, data *model.Category) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *MockCategoryRepo) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) GetAll(ctx context.Context) ([]*model.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) Delete(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Category), args.Error(1)
}

// --- Tests ---
func TestCategoryController_Create(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	ctrl := controller.NewCategoryController(mockRepo)

	expected := &model.Category{ID: 1, Name: "Electronics"}
	mockRepo.On("Create", mock.Anything, expected).Return(expected, nil)

	result, err := ctrl.Create(context.Background(), expected)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryController_Get_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	ctrl := controller.NewCategoryController(mockRepo)

	mockRepo.On("Get", mock.Anything, model.CategoryID(99)).Return(&model.Category{}, errors.New("not found"))

	_, err := ctrl.Get(context.Background(), 99)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
