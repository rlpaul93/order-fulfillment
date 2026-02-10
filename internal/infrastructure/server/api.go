package server

import (
	"net/http"

	"github.com/rlpaul93/order-fulfillment/internal/adapters/in"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// NewHandler sets up the HTTP routes and returns the handler
func NewHandler(prodSvc *service.ProductService, packSvc *service.PackService, fulfillSvc *service.PackFulfillmentService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", in.CreateProductHandler(prodSvc))
	mux.HandleFunc("/packs", in.CreatePackHandler(packSvc))
	mux.HandleFunc("/fulfill", in.PackFulfillmentHandler(fulfillSvc, packSvc))
	return mux
}
