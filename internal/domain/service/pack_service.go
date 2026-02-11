package service

import (
	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/port"
)

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
