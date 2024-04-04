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

func (handler *CheckpointHandler) UpdateCheckpoint(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	checkpointId := vars["checkpointId"]

	var checkpoint model.Checkpoint
	err := json.NewDecoder(request.Body).Decode(&checkpoint)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(checkpointId)
	if err != nil {
		println("error converting checkpointId to int: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CheckpointService.Update(&checkpoint, int32(id))
	if err != nil {
		println("Error while updating the checkpoint: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CheckpointHandler) GetCheckpoints(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tourId := vars["tourId"]
	tourID, err := strconv.Atoi(tourId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	checkpoints := handler.CheckpointService.GetCheckpoints(int32(tourID))
	json.NewEncoder(writer).Encode(checkpoints)
	writer.Header().Set("Content-Type", "application/json")
}
