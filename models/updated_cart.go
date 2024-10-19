package models

type UpdatedCart struct {
	Items         []CartItem `json:"items"`
	TotalPrice    float64    `json:"total_price"`
	TotalDiscount float64    `json:"total_discount"`
	FinalPrice    float64    `json:"final_price"`
}

type ApplicableCoupon struct {
	CouponID uint       `json:"coupon_id"`
	Type     CouponType `json:"type"`
	Discount float64    `json:"discount"`
}
