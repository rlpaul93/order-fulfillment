package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

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

// Handler for getting a pack by ID
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

// Handler for deleting a pack by ID
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

// Handler for listing packs by product ID
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
