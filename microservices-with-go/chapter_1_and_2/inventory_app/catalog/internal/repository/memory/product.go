package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"inventory.com/catalog/pkg/model"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

// Product handles in-memory storage for products.
type Product struct {
	mu    sync.RWMutex
	data  []*model.ProductBasic
	seqID int
}

// NewProduct returns a new in-memory Product repository.
func NewProduct() *Product {
	return &Product{
		data: make([]*model.ProductBasic, 0),
	}
}

// Create adds a new product to the in-memory store.
func (repo *Product) Create(ctx context.Context, input *model.ProductBasic) (*model.ProductBasic, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.seqID++
	input.ID = model.ProductID(repo.seqID)
	repo.data = append(repo.data, input)
	return input, nil
}

// Update modifies an existing product by ID.
func (repo *Product) Update(ctx context.Context, id model.ProductID, updated *model.ProductBasic) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, existing := repo.find(id)
	if existing == nil {
		return fmt.Errorf("%w: id=%d", ErrProductNotFound, id)
	}
	*existing = *updated
	existing.ID = id // Make sure ID doesn't get overwritten
	return nil
}

// Get retrieves a product by ID.
func (repo *Product) Get(ctx context.Context, id model.ProductID) (*model.ProductBasic, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	_, existing := repo.find(id)
	if existing == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrProductNotFound, id)
	}
	return existing, nil
}

// GetAll returns all products. Returns ErrProductNotFound if no entries exist.
func (repo *Product) GetAll(ctx context.Context) ([]*model.ProductBasic, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, ErrProductNotFound
	}
	return repo.data, nil
}

// Delete removes a product by ID and returns the deleted product.
func (repo *Product) Delete(ctx context.Context, id model.ProductID) (*model.ProductBasic, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	index, existing := repo.find(id)
	if existing == nil {
		return nil, fmt.Errorf("%w: id=%d", ErrProductNotFound, id)
	}
	repo.data = append(repo.data[:index], repo.data[index+1:]...)
	return existing, nil
}

// find locates a product by ID and returns index and pointer.
func (repo *Product) find(id model.ProductID) (int, *model.ProductBasic) {
	for i, p := range repo.data {
		if p.ID == id {
			return i, p
		}
	}
	return -1, nil
}
