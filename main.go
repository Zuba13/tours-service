package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tours-service.xws.com/handler"
	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
	"tours-service.xws.com/service"
)

func initDB() *gorm.DB {
	connectionStr := "user=postgres password=super dbname=tours host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil
	}

	db.AutoMigrate(&model.Equipment{})
	db.AutoMigrate(&model.Tour{})
	db.AutoMigrate(&model.Checkpoint{})
	db.AutoMigrate(&model.TourEquipment{})
	db.AutoMigrate(&model.TourExecution{})
	db.AutoMigrate(&model.CheckpointStatus{})

	fmt.Println("Successfully connected!")

	return db
}

func startServer(tourHandler *handler.TourHandler,
	checkpointHandler *handler.CheckpointHandler,
	equipmentHandler *handler.EquipmentHandler,
	tourExecutionHandler *handler.TourExecutionHandler) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create-tour", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/get-tours/{authorId}", tourHandler.GetAuthorTours).Methods("GET")
	router.HandleFunc("/add-checkpoint/{tourId}", checkpointHandler.CreateCheckpoint).Methods("POST")
	router.HandleFunc("/add-equipment/{tourId}", tourHandler.AddEquipment).Methods("POST")
	router.HandleFunc("/get-tour/{tourId}", tourHandler.GetTourById).Methods("GET")
	router.HandleFunc("/get-checkpoints/{tourId}", checkpointHandler.GetCheckpoints).Methods("GET")
	router.HandleFunc("/get-equipment", equipmentHandler.GetEquipment).Methods("GET")
	router.HandleFunc("/tourist/execute-tour/{tourId}/{touristId}", tourExecutionHandler.ExecuteTour).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	equipmentRepo := &repo.TourEquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	checkpointRepo := &repo.CheckpointRepository{DatabaseConnection: database}
	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}
	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourExecutionRepo := &repo.TourExecutionRepository{DatabaseConnection: database, CheckpointRepo: checkpointRepo}
	tourExecutionService := &service.TourExecutionService{TourExecutionRepo: tourExecutionRepo}
	tourExecutionHandler := &handler.TourExecutionHandler{TourExecutionService: tourExecutionService}

	startServer(tourHandler, checkpointHandler, equipmentHandler, tourExecutionHandler)

}
