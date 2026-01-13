package main

import (
	"log"
	"net/http"

	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	log.Println("Trip Service START")

	inMemRepository := repository.NewInMemRepository()
	svc := service.NewService(inMemRepository)
	httpHandler := h.NewHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
