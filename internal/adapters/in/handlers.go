package in

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// Handler for creating a product
func CreateProductHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p model.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			slog.Error("Failed to decode product", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := svc.Create(&p); err != nil {
			slog.Error("Failed to create product", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Product created", "product", p)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	}
}

// Handler for creating a pack
func CreatePackHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p model.Pack
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			slog.Error("Failed to decode pack", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := svc.Create(&p); err != nil {
			slog.Error("Failed to create pack", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Pack created", "pack", p)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	}
}

// Handler for pack fulfillment
func PackFulfillmentHandler(svc *service.PackFulfillmentService, packSvc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIDStr := r.URL.Query().Get("product_id")
		quantityStr := r.URL.Query().Get("quantity")
		slog.Info("Pack fulfillment request", "product_id", productIDStr, "quantity", quantityStr)
		productID, err := uuid.Parse(productIDStr)
		if err != nil {
			slog.Error("Invalid product_id UUID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			slog.Error("Invalid quantity", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		packs, err := packSvc.ListByProduct(productID)
		if err != nil || len(packs) == 0 {
			slog.Error("No packs found for product", "product_id", productIDStr)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var sizes []int
		for _, p := range packs {
			sizes = append(sizes, p.Size)
		}
		result := svc.FulfillOrder(quantity, sizes)
		slog.Info("Pack fulfillment result", "result", result)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
