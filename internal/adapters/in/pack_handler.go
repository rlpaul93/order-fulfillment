package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// ListPacksForProductHandler godoc
// @Summary List packs for a product
// @Description Get all packs for a specific product
// @Tags Products
// @Produce json
// @Param id path string true "Product UUID"
// @Success 200 {array} model.Pack
// @Failure 400 {string} string "Invalid product ID"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id}/packs [get]
func ListPacksForProductHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIDStr := r.PathValue("id")
		productID, err := uuid.Parse(productIDStr)
		if err != nil {
			slog.Error("Invalid product ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		packs, err := svc.ListByProduct(productID)
		if err != nil {
			slog.Error("Failed to list packs", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Packs listed", "product_id", productID, "count", len(packs))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(packs)
	}
}

// UpdatePacksForProductHandler godoc
// @Summary Update packs for a product
// @Description Replace all packs for a product with a new list of sizes
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product UUID"
// @Param sizes body []int true "Array of pack sizes"
// @Success 200 {array} model.Pack
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id}/packs [put]
func UpdatePacksForProductHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIDStr := r.PathValue("id")
		productID, err := uuid.Parse(productIDStr)
		if err != nil {
			slog.Error("Invalid product ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var sizes []int
		if err := json.NewDecoder(r.Body).Decode(&sizes); err != nil {
			slog.Error("Failed to decode sizes", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		packs, err := svc.ReplaceByProduct(productID, sizes)
		if err != nil {
			slog.Error("Failed to update packs", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Packs updated", "product_id", productID, "count", len(packs))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(packs)
	}
}
