package models

type Vehicle struct {
	Id_vehicle  uint   `gorm:"primaryKey" json:"id_vehicle"`
	VehicleName string `json:"vehicle_name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	Stock       string `json:"stock"`
	Image       string `json:"image"`
}

type Vehicles []Vehicles
