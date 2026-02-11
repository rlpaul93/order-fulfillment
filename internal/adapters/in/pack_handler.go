package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// CreatePackHandler godoc
// @Summary Create a new pack
// @Description Create a new pack for a product
// @Tags Packs
// @Accept json
// @Produce json
// @Param pack body model.Pack true "Pack to create"
// @Success 201 {object} model.Pack
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /packs [post]
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

// GetPackHandler godoc
// @Summary Get a pack by ID
// @Description Get a pack by its UUID
// @Tags Packs
// @Produce json
// @Param id path string true "Pack UUID"
// @Success 200 {object} model.Pack
// @Failure 400 {string} string "Invalid pack ID"
// @Failure 404 {string} string "Pack not found"
// @Router /packs/{id} [get]
func GetPackHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("Invalid pack ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p, err := svc.GetByID(id)
		if err != nil {
			slog.Error("Pack not found", "id", id, "error", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		slog.Info("Pack retrieved", "pack", p)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

// DeletePackHandler godoc
// @Summary Delete a pack by ID
// @Description Delete a pack by its UUID
// @Tags Packs
// @Param id path string true "Pack UUID"
// @Success 204 {string} string "Pack deleted"
// @Failure 400 {string} string "Invalid pack ID"
// @Failure 500 {string} string "Internal server error"
// @Router /packs/{id} [delete]
func DeletePackHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("Invalid pack ID", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := svc.Delete(id); err != nil {
			slog.Error("Failed to delete pack", "id", id, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Pack deleted", "id", id)
		w.WriteHeader(http.StatusNoContent)
	}
}

// ListPacksByProductHandler godoc
// @Summary List packs by product ID
// @Description Get a list of packs for a specific product
// @Tags Packs
// @Produce json
// @Param product_id query string true "Product UUID"
// @Success 200 {array} model.Pack
// @Failure 400 {string} string "Invalid product_id"
// @Failure 500 {string} string "Internal server error"
// @Router /packs [get]
func ListPacksByProductHandler(svc *service.PackService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIDStr := r.URL.Query().Get("product_id")
		productID, err := uuid.Parse(productIDStr)
		if err != nil {
			slog.Error("Invalid product_id UUID", "error", err)
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
