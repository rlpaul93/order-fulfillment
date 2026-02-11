package out

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
)

var ErrPackNotFound = errors.New("pack not found")

// PackRepositoryMem is an in-memory implementation of PackRepository.
type PackRepositoryMem struct {
	mu    sync.RWMutex
	packs map[uuid.UUID]*model.Pack
}

// NewPackRepositoryMem creates a new in-memory pack repository.
func NewPackRepositoryMem() *PackRepositoryMem {
	return &PackRepositoryMem{
		packs: make(map[uuid.UUID]*model.Pack),
	}
}

func (r *PackRepositoryMem) Create(pack *model.Pack) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	pack.ID = uuid.New()
	r.packs[pack.ID] = pack
	return nil
}

func (r *PackRepositoryMem) GetByID(id uuid.UUID) (*model.Pack, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.packs[id]
	if !ok {
		return nil, ErrPackNotFound
	}
	return p, nil
}

func (r *PackRepositoryMem) Update(pack *model.Pack) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.packs[pack.ID]; !ok {
		return ErrPackNotFound
	}
	r.packs[pack.ID] = pack
	return nil
}

func (r *PackRepositoryMem) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.packs[id]; !ok {
		return ErrPackNotFound
	}
	delete(r.packs, id)
	return nil
}

func (r *PackRepositoryMem) DeleteByProduct(productID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, pack := range r.packs {
		if pack.ProductID == productID {
			delete(r.packs, id)
		}
	}
	return nil
}

func (r *PackRepositoryMem) ListByProduct(productID uuid.UUID) ([]*model.Pack, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var packs []*model.Pack
	for _, pack := range r.packs {
		if pack.ProductID == productID {
			packs = append(packs, pack)
		}
	}
	return packs, nil
}
