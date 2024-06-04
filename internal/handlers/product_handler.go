package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yernazarius/pharmacy_store_golang/internal/domain/services"
	"github.com/yernazarius/pharmacy_store_golang/internal/infrastructure/messaging/nats"
)

type ProductHandler struct {
	Service    *services.ProductService
	NATSClient *nats.NATSClient
}

func NewProductHandler(db *sql.DB, natsClient *nats.NATSClient) *ProductHandler {
	repo := &repositories.ProductRepository{DB: db}
	service := &services.ProductService{Repo: repo}
	return &ProductHandler{Service: service, NATSClient: natsClient}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// Add other CRUD methods similarly...
