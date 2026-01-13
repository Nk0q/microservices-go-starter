package http

import (
	"encoding/json"
	"log"
	"net/http"

	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type Handler struct {
	service domain.TripService
}

func NewHandler(service domain.TripService) *Handler {
	return &Handler{
		service: service,
	}
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *Handler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Println("ddfsdsadsadsadd===============")

	t, err := s.service.GetRoute(r.Context(), &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Panic(err)
	}

	writeJSON(w, http.StatusCreated, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
