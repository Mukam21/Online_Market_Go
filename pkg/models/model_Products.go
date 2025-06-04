package models

type Product struct {
	ID          uint64  `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
