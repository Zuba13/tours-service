package model

type Checkpoint struct {
	Id          int32   `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	PictureURL  string  `json:"PictureURL"`
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
	TourId      int32   `json:"TourId"`
	Tour        *Tour   `json:"tour"`
}
