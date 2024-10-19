package strategies

import (
	"encoding/json"
	"errors"
	"math"

	"coupon-api/models"
)

type BxGyStrategy struct{}

type BxGyDetails struct {
	BuyProducts     []ProductQuantity `json:"buy_products"`
	GetProducts     []ProductQuantity `json:"get_products"`
	RepetitionLimit uint              `json:"repetition_limit"`
}

type ProductQuantity struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func (s *BxGyStrategy) CalculateDiscount(coupon *models.Coupon, cart *models.Cart) (float64, error) {
	var details BxGyDetails
	data, err := json.Marshal(coupon.Details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return 0, errors.New("invalid coupon details")
	}

	timesApplicable := s.calculateTimesApplicable(details, cart)
	if timesApplicable == 0 {
		return 0, nil
	}

	totalDiscount := s.calculateTotalDiscount(details, cart, timesApplicable)
	return totalDiscount, nil
}

func (s *BxGyStrategy) ApplyCoupon(coupon *models.Coupon, cart *models.Cart) (*models.UpdatedCart, error) {
	var details BxGyDetails
	data, err := json.Marshal(coupon.Details)
	if err != nil {
		return nil, errors.New("invalid coupon details")
	}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return nil, errors.New("invalid coupon details")
	}

	timesApplicable := s.calculateTimesApplicable(details, cart)
	if timesApplicable == 0 {
		return nil, errors.New("coupon conditions not met")
	}

	if details.RepetitionLimit > 0 && timesApplicable > details.RepetitionLimit {
		timesApplicable = details.RepetitionLimit
	}

	updatedItems := make([]models.CartItem, len(cart.Items))
	copy(updatedItems, cart.Items)

	totalDiscount := s.applyDiscountToCart(details, updatedItems, timesApplicable)

	updatedCart := &models.UpdatedCart{
		Items:         updatedItems,
		TotalPrice:    calculateCartTotal(cart),
		TotalDiscount: totalDiscount,
		FinalPrice:    calculateCartTotal(cart) - totalDiscount,
	}
	return updatedCart, nil
}

func (s *BxGyStrategy) calculateTimesApplicable(details BxGyDetails, cart *models.Cart) uint {
	minTimes := uint(math.MaxUint32)
	for _, bp := range details.BuyProducts {
		quantityInCart := s.getQuantityInCart(bp.ProductID, cart)
		if bp.Quantity == 0 {
			return 0
		}
		times := quantityInCart / bp.Quantity
		if times < minTimes {
			minTimes = times
		}
	}
	return minTimes
}

func (s *BxGyStrategy) getQuantityInCart(productID uint, cart *models.Cart) uint {
	for _, item := range cart.Items {
		if item.ProductID == productID {
			return item.Quantity
		}
	}
	return 0
}

func (s *BxGyStrategy) calculateTotalDiscount(details BxGyDetails, cart *models.Cart, timesApplicable uint) float64 {
	totalDiscount := 0.0
	for _, gp := range details.GetProducts {
		quantityNeeded := gp.Quantity * timesApplicable
		quantityInCart := s.getQuantityInCart(gp.ProductID, cart)
		quantityToDiscount := min(quantityNeeded, quantityInCart)
		price := s.getPriceOfProduct(gp.ProductID, cart)
		totalDiscount += float64(quantityToDiscount) * price
	}
	return totalDiscount
}

func (s *BxGyStrategy) applyDiscountToCart(details BxGyDetails, items []models.CartItem, timesApplicable uint) float64 {
	totalDiscount := 0.0
	for i, item := range items {
		for _, gp := range details.GetProducts {
			if item.ProductID == gp.ProductID {
				quantityNeeded := gp.Quantity * timesApplicable
				quantityToDiscount := min(quantityNeeded, item.Quantity)
				itemDiscount := float64(quantityToDiscount) * item.Price
				items[i].TotalDiscount += itemDiscount
				totalDiscount += itemDiscount
			}
		}
	}
	return totalDiscount
}

func (s *BxGyStrategy) getPriceOfProduct(productID uint, cart *models.Cart) float64 {
	for _, item := range cart.Items {
		if item.ProductID == productID {
			return item.Price
		}
	}
	return 0.0
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
