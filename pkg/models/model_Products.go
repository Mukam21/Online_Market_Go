package models

type Product struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	Name        string
	Description string
	Price       float64
	Quantity    int
}
