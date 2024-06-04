package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yernazarius/pharmacy_store_golang/internal/domain/services"
	"github.com/yernazarius/pharmacy_store_golang/internal/infrastructure/messaging/nats"
)

type UserHandler struct {
	Service    *services.UserService
	NATSClient *nats.NATSClient
}

func NewUserHandler(db *sql.DB, natsClient *nats.NATSClient) *UserHandler {
	repo := &repositories.UserRepository{DB: db}
	service := &services.UserService{Repo: repo}
	return &UserHandler{Service: service, NATSClient: natsClient}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Add other CRUD methods similarly...
