package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// CreateProductHandler godoc
// @Summary Create a new product
// @Description Create a new product with a name
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product to create"
// @Success 201 {object} model.Product
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /products [post]
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

// GetProductHandler godoc
// @Summary Get a product by ID
// @Description Get a product by its UUID
// @Tags Products
// @Produce json
// @Param id path string true "Product UUID"
// @Success 200 {object} model.Product
// @Failure 400 {string} string "Invalid product ID"
// @Failure 404 {string} string "Product not found"
// @Router /products/{id} [get]
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

// DeleteProductHandler godoc
// @Summary Delete a product by ID
// @Description Delete a product by its UUID
// @Tags Products
// @Param id path string true "Product UUID"
// @Success 204 {string} string "Product deleted"
// @Failure 400 {string} string "Invalid product ID"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id} [delete]
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

// ListProductsHandler godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {string} string "Internal server error"
// @Router /products [get]
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
