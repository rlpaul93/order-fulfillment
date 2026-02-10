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

func (s *ProductService) GetByID(id int64) (*model.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) Update(product *model.Product) error {
	return s.Repo.Update(product)
}

func (s *ProductService) Delete(id int64) error {
	return s.Repo.Delete(id)
}

func (s *ProductService) List() ([]*model.Product, error) {
	return s.Repo.List()
}

// PackService provides business logic for packs.
type PackService struct {
	Repo port.PackRepository
}

func (s *PackService) Create(pack *model.Pack) error {
	return s.Repo.Create(pack)
}

func (s *PackService) GetByID(id uuid.UUID) (*model.Pack, error) {
	return s.Repo.GetByID(id)
}

func (s *PackService) Update(pack *model.Pack) error {
	return s.Repo.Update(pack)
}

func (s *PackService) Delete(id uuid.UUID) error {
	return s.Repo.Delete(id)
}

func (s *PackService) ListByProduct(productID uuid.UUID) ([]*model.Pack, error) {
	return s.Repo.ListByProduct(productID)
}
