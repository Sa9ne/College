package models

import "time"

type Users struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Wallet   int    `json:"wallet"`
}

type Car struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Mileage int    `json:"mileage"`
	Year    int    `json:"Year"`
}

type Orders struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CarID     uint      `json:"car_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"bought_at"`

	User Users `gorm:"foreignKey:UserID"`
	Car  Car   `gorm:"foreignKey:CarID"`
}
