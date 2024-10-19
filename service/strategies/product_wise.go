package strategies

import (
	"encoding/json"
	"errors"

	"coupon-api/models"
)

type ProductWiseStrategy struct{}

type ProductWiseDetails struct {
	ProductID uint    `json:"product_id"`
	Discount  float64 `json:"discount"`
}

func (s *ProductWiseStrategy) CalculateDiscount(coupon *models.Coupon, cart *models.Cart) (float64, error) {
	var details ProductWiseDetails
	data, err := json.Marshal(coupon.Details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}

	totalDiscount := 0.0
	for _, item := range cart.Items {
		if item.ProductID == details.ProductID {
			totalDiscount += item.Price * float64(item.Quantity) * (details.Discount / 100)
		}
	}
	return totalDiscount, nil
}

func (s *ProductWiseStrategy) ApplyCoupon(coupon *models.Coupon, cart *models.Cart) (*models.UpdatedCart, error) {
	var details ProductWiseDetails
	data, err := json.Marshal(coupon.Details)
	if err != nil {
		return nil, errors.New("invalid coupon details")
	}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return nil, errors.New("invalid coupon details")
	}

	totalDiscount := 0.0
	updatedItems := make([]models.CartItem, len(cart.Items))
	copy(updatedItems, cart.Items)

	for i, item := range updatedItems {
		if item.ProductID == details.ProductID {
			itemDiscount := item.Price * float64(item.Quantity) * (details.Discount / 100)
			updatedItems[i].TotalDiscount = itemDiscount
			totalDiscount += itemDiscount
		}
	}

	updatedCart := &models.UpdatedCart{
		Items:         updatedItems,
		TotalPrice:    calculateCartTotal(cart),
		TotalDiscount: totalDiscount,
		FinalPrice:    calculateCartTotal(cart) - totalDiscount,
	}
	return updatedCart, nil
}
