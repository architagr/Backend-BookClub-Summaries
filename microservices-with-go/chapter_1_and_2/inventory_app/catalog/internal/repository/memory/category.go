package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"inventory.com/catalog/pkg/model"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

// Category represents an in-memory repository for categories.
type Category struct {
	mu    sync.RWMutex
	data  []*model.Category
	seqID int
}

// NewCategory returns a new in-memory Category repository.
func NewCategory() *Category {
	return &Category{
		data:  make([]*model.Category, 0),
		seqID: 0,
	}
}

// Create adds a new category to the in-memory store.
func (repo *Category) Create(ctx context.Context, data *model.Category) (*model.Category, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.seqID++
	data.ID = model.CategoryID(repo.seqID)
	repo.data = append(repo.data, data)

	return data, nil
}

// Update updates an existing category. Returns ErrCategoryNotFound if not found.
func (repo *Category) Update(ctx context.Context, id model.CategoryID, data *model.Category) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, existing := repo.find(id)
	if existing == nil {
		return fmt.Errorf("%w: id=%d", ErrCategoryNotFound, id)
	}
	existing.Name = data.Name
	return nil
}

// GetAll returns all categories. Returns ErrCategoryNotFound if no categories exist.
// TODO: Add pagination support
func (repo *Category) GetAll(ctx context.Context) ([]*model.Category, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, ErrCategoryNotFound
	}
	return repo.data, nil
}

// Delete removes a category by ID. Returns the deleted category or ErrCategoryNotFound.
func (repo *Category) Delete(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	index, found := repo.find(id)
	if found == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrCategoryNotFound, id)
	}

	repo.data = append(repo.data[:index], repo.data[index+1:]...)
	return found, nil
}

// Get retrieves a category by ID. Returns ErrCategoryNotFound if not found.
func (repo *Category) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	_, found := repo.find(id)
	if found == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrCategoryNotFound, id)
	}
	return found, nil
}

// find returns the index and pointer to the category, or -1 and nil if not found.
func (repo *Category) find(id model.CategoryID) (int, *model.Category) {
	for i, d := range repo.data {
		if d.ID == id {
			return i, d
		}
	}
	return -1, nil
}
