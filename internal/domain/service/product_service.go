package service

import (
	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/port"
)

// ProductService provides business logic for products.
type ProductService struct {
	Repo port.ProductRepository
}

func (s *ProductService) Create(product *model.Product) error {
	return s.Repo.Create(product)
}

func (s *ProductService) GetByID(id uuid.UUID) (*model.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) Update(product *model.Product) error {
	return s.Repo.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.Repo.Delete(id)
}

func (s *ProductService) List() ([]*model.Product, error) {
	return s.Repo.List()
}
