package http

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/yernazarius/pharmacy_store_golang/internal/handlers"
	"github.com/yernazarius/pharmacy_store_golang/internal/infrastructure/messaging/nats"
)

func NewRouter(db *sql.DB, natsClient *nats.NATSClient) *mux.Router {
	router := mux.NewRouter()

	productHandler := handlers.NewProductHandler(db, natsClient)
	userHandler := handlers.NewUserHandler(db, natsClient)

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return router
}
