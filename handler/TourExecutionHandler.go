package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tours-service.xws.com/service"
)

type TourExecutionHandler struct {
	TourExecutionService *service.TourExecutionService
}

func (handler *TourExecutionHandler) ExecuteTour(w http.ResponseWriter, r *http.Request) {
	tourId, err := strconv.Atoi(mux.Vars(r)["tourId"])
	if err != nil {
		http.Error(w, "Invalid tour id", http.StatusBadRequest)
		return
	}
	touristId, err := strconv.Atoi(mux.Vars(r)["touristId"])
	if err != nil {
		http.Error(w, "Invalid tourist id", http.StatusBadRequest)
		return
	}

	tourExecution, err := handler.TourExecutionService.ExecuteTour(int32(tourId), int32(touristId))
	if err != nil {
		http.Error(w, "Error executing tour: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Optionally, you can return the created tour execution object in the response
	responseJSON, err := json.Marshal(tourExecution)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
