package model

type Equipment struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TourEquipment struct {
	TourId      int32 `json:"tour_id"`
	EquipmentId int32 `json:"equipment_id"`
}
