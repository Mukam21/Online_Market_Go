package models

type OrderRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
