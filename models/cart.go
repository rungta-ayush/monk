package models

type Cart struct {
	UserID uint       `json:"user_id,omitempty"`
	Items  []CartItem `json:"items" binding:"required,dive"`
}

type CartItem struct {
	ProductID     uint    `json:"product_id" binding:"required"`
	Quantity      uint    `json:"quantity" binding:"required,min=1"`
	Price         float64 `json:"price" binding:"required,gt=0"`
	TotalDiscount float64 `json:"total_discount,omitempty"`
}
