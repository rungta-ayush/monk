package strategies

import (
	"encoding/json"
	"errors"

	"coupon-api/models"
)

type CartWiseStrategy struct{}

type CartWiseDetails struct {
	Threshold float64 `json:"threshold"`
	Discount  float64 `json:"discount"`
}

func (s *CartWiseStrategy) CalculateDiscount(coupon *models.Coupon, cart *models.Cart) (float64, error) {
	var details CartWiseDetails
	data, err := json.Marshal(coupon.Details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}

	totalAmount := calculateCartTotal(cart)
	if totalAmount > details.Threshold {
		discount := totalAmount * (details.Discount / 100)
		return discount, nil
	}
	return 0, nil
}

func (s *CartWiseStrategy) ApplyCoupon(coupon *models.Coupon, cart *models.Cart) (*models.UpdatedCart, error) {
	discount, err := s.CalculateDiscount(coupon, cart)
	if err != nil {
		return nil, err
	}

	updatedCart := &models.UpdatedCart{
		Items:         cart.Items,
		TotalPrice:    calculateCartTotal(cart),
		TotalDiscount: discount,
		FinalPrice:    calculateCartTotal(cart) - discount,
	}
	return updatedCart, nil
}

func calculateCartTotal(cart *models.Cart) float64 {
	total := 0.0
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
