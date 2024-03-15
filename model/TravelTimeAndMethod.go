package model

type TravelMethod int

const (
	CAR Difficult = iota
	BICYCLE
	WALKING
)

type TravelTimeAndMethod struct {
	TravelTime int32        `json:"travelTime"`
	Method     TravelMethod `json:"travelMethod"`
}
