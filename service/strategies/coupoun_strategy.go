package strategies

import (
	"coupon-api/models"
)

type CouponStrategy interface {
	CalculateDiscount(coupon *models.Coupon, cart *models.Cart) (float64, error)
	ApplyCoupon(coupon *models.Coupon, cart *models.Cart) (*models.UpdatedCart, error)
}

type CouponStrategyFactory interface {
	GetStrategy(couponType models.CouponType) CouponStrategy
}

type strategyFactory struct {
	strategies map[models.CouponType]CouponStrategy
}

func NewCouponStrategyFactory() CouponStrategyFactory {
	factory := &strategyFactory{
		strategies: make(map[models.CouponType]CouponStrategy),
	}
	factory.strategies[models.CartWise] = &CartWiseStrategy{}
	factory.strategies[models.ProductWise] = &ProductWiseStrategy{}
	factory.strategies[models.BxGy] = &BxGyStrategy{}
	// Additional strategies can be registered here
	return factory
}

func (f *strategyFactory) GetStrategy(couponType models.CouponType) CouponStrategy {
	return f.strategies[couponType]
}
