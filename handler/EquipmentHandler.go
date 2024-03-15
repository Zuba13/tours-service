package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tours-service.xws.com/model"
	"tours-service.xws.com/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) SaveTourEquipment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tourIDStr := vars["tourId"]
	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		fmt.Println("error parsing tour ID:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var equipment []model.Equipment
	if err := json.NewDecoder(request.Body).Decode(&equipment); err != nil {
		fmt.Println("error parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var equipmentIDs []int32
	for _, equip := range equipment {
		equipmentIDs = append(equipmentIDs, equip.Id)
	}

	if err := handler.EquipmentService.SaveTourEquipment(equipment, int32(tourID)); err != nil {
		fmt.Println("error saving tour equipment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (handler *EquipmentHandler) GetEquipment(writer http.ResponseWriter, request *http.Request) {
	equipment := handler.EquipmentService.GetEquipment()
	json.NewEncoder(writer).Encode(equipment)
	writer.Header().Set("Content-Type", "application/json")
}
