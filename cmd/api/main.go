package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rlpaul93/order-fulfillment/cmd/api/config"
	"github.com/rlpaul93/order-fulfillment/cmd/api/factory"
	"github.com/rlpaul93/order-fulfillment/docs"
	"github.com/rlpaul93/order-fulfillment/internal/infrastructure/db"
	"github.com/rlpaul93/order-fulfillment/internal/infrastructure/server"
)

// @title Order Fulfillment API
// @version 1.0
// @description REST API for product and pack management with optimal order fulfillment
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	cfg := config.Load()

	// Set Swagger host dynamically
	docs.SwaggerInfo.Host = cfg.SwaggerHost
	log.Printf("Swagger host: %s", cfg.SwaggerHost)

	var dbConn *sql.DB
	if cfg.StorageMode == "postgres" {
		var err error
		dbConn, err = db.NewConnection(cfg.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close()
		log.Println("Using PostgreSQL storage")
	} else {
		log.Println("Using in-memory storage")
	}

	prodSvc, packSvc, fulfillSvc := factory.BuildServices(dbConn)
	handler := server.NewHandler(prodSvc, packSvc, fulfillSvc)

	log.Printf("API running on :%s", cfg.APIPort)
	log.Fatal(http.ListenAndServe(":"+cfg.APIPort, handler))
}
