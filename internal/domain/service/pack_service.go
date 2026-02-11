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

// ReplaceByProduct deletes all existing packs for a product and creates new ones with the given sizes.
func (s *PackService) ReplaceByProduct(productID uuid.UUID, sizes []int) ([]*model.Pack, error) {
	if err := s.Repo.DeleteByProduct(productID); err != nil {
		return nil, err
	}
	var packs []*model.Pack
	for _, size := range sizes {
		pack := &model.Pack{ProductID: productID, Size: size}
		if err := s.Repo.Create(pack); err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}
	return packs, nil
}
