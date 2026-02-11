package in

import (
	"encoding/json"
	"log/slog"
	"net/http"

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
