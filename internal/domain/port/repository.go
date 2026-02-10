package port

import (
	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
)

// ProductRepository defines CRUD operations for products.
type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id int64) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id int64) error
	List() ([]*model.Product, error)
}

// PackRepository defines CRUD operations for packs.
type PackRepository interface {
	Create(pack *model.Pack) error
	GetByID(id uuid.UUID) (*model.Pack, error)
	Update(pack *model.Pack) error
	Delete(id uuid.UUID) error
	ListByProduct(productID uuid.UUID) ([]*model.Pack, error)
}
