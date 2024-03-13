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

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(request.Body).Decode(&tour)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tour.Status = model.DRAFT
	tour.Price = 0
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) GetAuthorTours(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	authorId := vars["authorId"]
	fmt.Println("Usao sam u metodu; authorId: ", authorId)
	if authorId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	authorID, err := strconv.Atoi(authorId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tours := handler.TourService.GetAuthorTours(int32(authorID))
	json.NewEncoder(writer).Encode(tours)
	writer.Header().Set("Content-Type", "application/json")
}
