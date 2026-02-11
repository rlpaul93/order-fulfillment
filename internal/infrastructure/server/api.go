package server

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/rlpaul93/order-fulfillment/internal/adapters/in"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// NewHandler sets up the HTTP routes and returns the handler
func NewHandler(prodSvc *service.ProductService, packSvc *service.PackService, fulfillSvc *service.PackFulfillmentService) http.Handler {
	mux := http.NewServeMux()

	// Swagger UI
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	// Product routes
	mux.HandleFunc("POST /products", in.CreateProductHandler(prodSvc))
	mux.HandleFunc("GET /products", in.ListProductsHandler(prodSvc))
	mux.HandleFunc("GET /products/{id}", in.GetProductHandler(prodSvc))
	mux.HandleFunc("DELETE /products/{id}", in.DeleteProductHandler(prodSvc))

	// Pack routes (nested under products)
	mux.HandleFunc("GET /products/{id}/packs", in.ListPacksForProductHandler(packSvc))
	mux.HandleFunc("PUT /products/{id}/packs", in.UpdatePacksForProductHandler(packSvc))

	// Fulfillment route
	mux.HandleFunc("GET /fulfill", in.PackFulfillmentHandler(fulfillSvc, packSvc))

	return mux
}
