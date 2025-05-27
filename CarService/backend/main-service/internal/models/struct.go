package models

import "time"

type Client struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthdate time.Time `gorm:"type:date"`
	Phone     string    `json:"phone"`
}

type Car struct {
	VIN          int    `json:"vin"`
	Brand        string `json:"brand"`
	Name         string `json:"name"`
	Power        int    `json:"power"`
	Color        string `json:"color"`
	Price        int    `json:"price"`
	Mileage      int    `json:"mileage"`
	Year         int    `json:"year"`
	Condition    string `json:"condition"`
	Owners       int    `json:"owner"`
	BoughtStatus bool   `json:"boughtStatus"`
	Img          string `json:"img"`
}
