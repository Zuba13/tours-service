package model

type Checkpoint struct {
	Id          int32   `json:"id"`
	TourId      int32   `json:"tour_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Picture     string  `json:"picture"`
	Tour        *Tour   `json:"tour"`
}
