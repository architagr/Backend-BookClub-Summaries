package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"inventory.com/catalog/pkg/model"
)

var (
	ErrSubCategoryNotFound = errors.New("sub-category not found")
)

// SubCategory handles in-memory storage for sub-categories.
type SubCategory struct {
	mu    sync.RWMutex
	data  []*model.SubCategoryBasic
	seqID int
}

// NewSubCategory returns a new in-memory SubCategory repository.
func NewSubCategory() *SubCategory {
	return &SubCategory{
		data: make([]*model.SubCategoryBasic, 0),
	}
}

// Create adds a new sub-category to the in-memory store.
func (repo *SubCategory) Create(ctx context.Context, input *model.SubCategoryBasic) (*model.SubCategoryBasic, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.seqID++
	input.BaseInfo.ID = model.SubCategoryID(repo.seqID)
	repo.data = append(repo.data, input)
	return input, nil
}

// Update modifies an existing sub-category by ID.
func (repo *SubCategory) Update(ctx context.Context, id model.SubCategoryID, updated *model.SubCategoryBasic) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, existing := repo.find(id)
	if existing == nil {
		return fmt.Errorf("%w: id=%d", ErrSubCategoryNotFound, id)
	}
	existing.BaseInfo.Name = updated.BaseInfo.Name
	existing.CatID = updated.CatID
	return nil
}

// Get returns a sub-category by ID.
func (repo *SubCategory) Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	_, existing := repo.find(id)
	if existing == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrSubCategoryNotFound, id)
	}
	return existing, nil
}

// GetAll returns all sub-categories.
// Returns ErrSubCategoryNotFound if store is empty.
func (repo *SubCategory) GetAll(ctx context.Context) ([]*model.SubCategoryBasic, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, ErrSubCategoryNotFound
	}
	return repo.data, nil
}

// Delete removes a sub-category by ID and returns the deleted item.
func (repo *SubCategory) Delete(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	index, existing := repo.find(id)
	if existing == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrSubCategoryNotFound, id)
	}
	repo.data = append(repo.data[:index], repo.data[index+1:]...)
	return existing, nil
}

// find locates the sub-category by ID and returns index and pointer.
func (repo *SubCategory) find(id model.SubCategoryID) (int, *model.SubCategoryBasic) {
	for i, d := range repo.data {
		if d.BaseInfo.ID == id {
			return i, d
		}
	}
	return -1, nil
}
