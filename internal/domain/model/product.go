package model

import "github.com/google/uuid"

// Product represents a product with customizable packs.
type Product struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Pack represents a pack size for a product.
type Pack struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Size      int       `json:"size"`
}
