package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

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
