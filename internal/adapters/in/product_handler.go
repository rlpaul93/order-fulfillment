package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

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

// Handler for getting a product by ID
func GetProductHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("Invalid product ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p, err := svc.GetByID(id)
		if err != nil {
			slog.Error("Product not found", "id", id, "error", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		slog.Info("Product retrieved", "product", p)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

// Handler for deleting a product by ID
func DeleteProductHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("Invalid product ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := svc.Delete(id); err != nil {
			slog.Error("Failed to delete product", "id", id, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Product deleted", "id", id)
		w.WriteHeader(http.StatusNoContent)
	}
}

// Handler for listing all products
func ListProductsHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := svc.List()
		if err != nil {
			slog.Error("Failed to list products", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Products listed", "count", len(products))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}
