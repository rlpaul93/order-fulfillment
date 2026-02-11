package in

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

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
