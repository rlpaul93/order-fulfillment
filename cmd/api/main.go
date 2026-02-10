package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rlpaul93/order-fulfillment/cmd/api/config"
	"github.com/rlpaul93/order-fulfillment/cmd/api/factory"
	"github.com/rlpaul93/order-fulfillment/internal/infrastructure/server"
)

func main() {
	cfg := config.Load()
	dbConn, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	prodSvc, packSvc, fulfillSvc := factory.BuildServices(dbConn)
	handler := server.NewHandler(prodSvc, packSvc, fulfillSvc)

	log.Printf("API running on :%s", cfg.APIPort)
	log.Fatal(http.ListenAndServe(":"+cfg.APIPort, handler))
}
