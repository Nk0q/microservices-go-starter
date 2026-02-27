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

	ctx := r.Context()

	osrmResp, err := s.service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Printf("failed to get route from OSRM: %v", err)
		http.Error(w, "failed to get route", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusOK, osrmResp.ToTripPreview()); err != nil {
		log.Printf("failed to write json response: %v", err)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
