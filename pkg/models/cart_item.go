package models

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}
