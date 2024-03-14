package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tours-service.xws.com/model"
	"tours-service.xws.com/service"
)

type CheckpointHandler struct {
	CheckpointService *service.CheckpointService
}

func (handler *CheckpointHandler) CreateCheckpoint(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tourId := vars["tourId"]
	var checkpoint model.Checkpoint
	err := json.NewDecoder(request.Body).Decode(&checkpoint)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tourID, err := strconv.Atoi(tourId)
	err = handler.CheckpointService.Create(&checkpoint, int32(tourID))
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
