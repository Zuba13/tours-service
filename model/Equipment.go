package model

type Equipment struct {
	Id          int32  `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type TourEquipment struct {
	TourId      int32 `json:"TourId"`
	EquipmentId int32 `json:"EquipmentId"`
}
