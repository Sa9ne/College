package models

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
