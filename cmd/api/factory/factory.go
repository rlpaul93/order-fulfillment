package factory

import (
	"database/sql"

	"github.com/rlpaul93/order-fulfillment/internal/adapters/out"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// BuildServices wires up dependencies for the API
func BuildServices(dbConn *sql.DB) (prodSvc *service.ProductService, packSvc *service.PackService, fulfillSvc *service.PackFulfillmentService) {
	prodRepo := &out.ProductRepositoryPg{DB: dbConn}
	packRepo := &out.PackRepositoryPg{DB: dbConn}
	prodSvc = &service.ProductService{Repo: prodRepo}
	packSvc = &service.PackService{Repo: packRepo}
	fulfillSvc = &service.PackFulfillmentService{}
	return
}
