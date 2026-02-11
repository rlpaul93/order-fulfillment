package factory

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/adapters/out"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/port"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// BuildServices wires up dependencies for the API
// If dbConn is nil, in-memory repositories are used
func BuildServices(dbConn *sql.DB) (prodSvc *service.ProductService, packSvc *service.PackService, fulfillSvc *service.PackFulfillmentService) {
	var prodRepo port.ProductRepository
	var packRepo port.PackRepository

	if dbConn != nil {
		prodRepo = &out.ProductRepositoryPg{DB: dbConn}
		packRepo = &out.PackRepositoryPg{DB: dbConn}
	} else {
		prodRepo = out.NewProductRepositoryMem()
		packRepo = out.NewPackRepositoryMem()
		seedDefaultData(prodRepo, packRepo)
	}

	prodSvc = &service.ProductService{Repo: prodRepo}
	packSvc = &service.PackService{Repo: packRepo}
	fulfillSvc = &service.PackFulfillmentService{}
	return
}

// seedDefaultData adds a default product with packs for in-memory storage
func seedDefaultData(prodRepo port.ProductRepository, packRepo port.PackRepository) {
	product := &model.Product{
		ID:   uuid.New(),
		Name: "Default Product",
	}
	if err := prodRepo.Create(product); err != nil {
		log.Printf("Failed to seed default product: %v", err)
		return
	}

	packSizes := []int{250, 500, 1000, 2000, 5000}
	for _, size := range packSizes {
		pack := &model.Pack{
			ProductID: product.ID,
			Size:      size,
		}
		if err := packRepo.Create(pack); err != nil {
			log.Printf("Failed to seed pack size %d: %v", size, err)
		}
	}

	log.Printf("Seeded default product '%s' with packs: %v", product.Name, packSizes)
}
