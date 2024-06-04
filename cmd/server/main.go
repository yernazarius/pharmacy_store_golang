package main

import (
	"log"
	"net/http"

	"github.com/yourusername/yourprojectname/configs"
	"github.com/yourusername/yourprojectname/internal/infrastructure/db/postgres"
	"github.com/yourusername/yourprojectname/internal/infrastructure/http"
	"github.com/yourusername/yourprojectname/internal/infrastructure/messaging/nats"
	"github.com/yourusername/yourprojectname/internal/middleware"
)

func main() {
	config := configs.LoadConfig()

	db, err := postgres.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	natsClient := nats.NewNATSClient(config)
	defer natsClient.Close()

	router := http.NewRouter(db, natsClient)
	router.Use(middleware.Logger)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
