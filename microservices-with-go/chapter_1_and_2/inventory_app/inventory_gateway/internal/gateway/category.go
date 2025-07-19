package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"inventory.com/catalog/pkg/model"
)

var (
	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = fmt.Errorf("resource not found")
)

// Gateway defines a movie metadata HTTP gateway.
type CategoryGateway struct {
	addr string
}

// NewCategoryGateway creates a new HTTP gateway for a movie metadata service.
func NewCategoryGateway(addr string) *CategoryGateway {
	return &CategoryGateway{addr}
}
func (g *CategoryGateway) Create(ctx context.Context, data *model.Category) (*model.Category, error) {
	// Create a new HTTP request to the category creation endpoint
	// with the provided context.
	// The request body should contain the JSON representation of the category data.
	// The response should be unmarshalled into a model.Category object.
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, g.addr+"/categories", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (g *CategoryGateway) Update(ctx context.Context, id model.CategoryID, data *model.Category) (*model.Category, error) {

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/categories/%d", g.addr, int(id)), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (g *CategoryGateway) Get(ctx context.Context, id model.CategoryID) (*model.Category, error) {
	var data *model.Category
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories/%d", g.addr, int(id)), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// func (g *CategoryGateway) GetAll(ctx context.Context) ([]*model.Category, error)                    {}
// func (g *CategoryGateway) Delete(ctx context.Context, id model.CategoryID) (*model.Category, error) {}
