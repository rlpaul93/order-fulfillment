package out

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
)

var ErrProductNotFound = errors.New("product not found")

// ProductRepositoryMem is an in-memory implementation of ProductRepository.
type ProductRepositoryMem struct {
	mu       sync.RWMutex
	products map[uuid.UUID]*model.Product
}

// NewProductRepositoryMem creates a new in-memory product repository.
func NewProductRepositoryMem() *ProductRepositoryMem {
	return &ProductRepositoryMem{
		products: make(map[uuid.UUID]*model.Product),
	}
}

func (r *ProductRepositoryMem) Create(product *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	product.ID = uuid.New()
	r.products[product.ID] = product
	return nil
}

func (r *ProductRepositoryMem) GetByID(id uuid.UUID) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.products[id]
	if !ok {
		return nil, ErrProductNotFound
	}
	return p, nil
}

func (r *ProductRepositoryMem) Update(product *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.products[product.ID]; !ok {
		return ErrProductNotFound
	}
	r.products[product.ID] = product
	return nil
}

func (r *ProductRepositoryMem) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.products[id]; !ok {
		return ErrProductNotFound
	}
	delete(r.products, id)
	return nil
}

func (r *ProductRepositoryMem) List() ([]*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	products := make([]*model.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}
	return products, nil
}
