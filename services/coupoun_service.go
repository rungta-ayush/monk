package services

import (
	"errors"
	"time"

	"coupon-api/service/strategies"

	"coupon-api/models"
	"coupon-api/repositories"
)

type CouponService interface {
	CreateCoupon(coupon *models.Coupon) error
	GetAllCoupons() ([]models.Coupon, error)
	GetCouponByID(id uint) (*models.Coupon, error)
	UpdateCoupon(coupon *models.Coupon) error
	DeleteCoupon(id uint) error
	GetApplicableCoupons(cart *models.Cart) ([]models.ApplicableCoupon, error)
	ApplyCoupon(couponID uint, cart *models.Cart) (*models.UpdatedCart, error)
}

type couponService struct {
	repo            repositories.CouponRepository
	strategyFactory strategies.CouponStrategyFactory
}

func NewCouponService(repo repositories.CouponRepository, factory strategies.CouponStrategyFactory) CouponService {
	return &couponService{
		repo:            repo,
		strategyFactory: factory,
	}
}

func (s *couponService) CreateCoupon(coupon *models.Coupon) error {
	return s.repo.CreateCoupon(coupon)
}

func (s *couponService) GetAllCoupons() ([]models.Coupon, error) {
	return s.repo.GetAllCoupons()
}

func (s *couponService) GetCouponByID(id uint) (*models.Coupon, error) {
	return s.repo.GetCouponByID(id)
}

func (s *couponService) UpdateCoupon(coupon *models.Coupon) error {
	return s.repo.UpdateCoupon(coupon)
}

func (s *couponService) DeleteCoupon(id uint) error {
	return s.repo.DeleteCoupon(id)
}

func (s *couponService) GetApplicableCoupons(cart *models.Cart) ([]models.ApplicableCoupon, error) {
	coupons, err := s.repo.GetAllCoupons()
	if err != nil {
		return nil, err
	}

	applicableCoupons := []models.ApplicableCoupon{}
	for _, coupon := range coupons {
		if !s.isCouponApplicable(&coupon, cart) {
			continue
		}

		strategy := s.strategyFactory.GetStrategy(coupon.Type)
		if strategy == nil {
			continue
		}

		discount, err := strategy.CalculateDiscount(&coupon, cart)
		if err != nil {
			continue
		}

		if discount > 0 {
			applicableCoupons = append(applicableCoupons, models.ApplicableCoupon{
				CouponID: coupon.ID,
				Type:     coupon.Type,
				Discount: discount,
			})
		}
	}
	return applicableCoupons, nil
}

func (s *couponService) ApplyCoupon(couponID uint, cart *models.Cart) (*models.UpdatedCart, error) {
	coupon, err := s.repo.GetCouponByID(couponID)
	if err != nil {
		return nil, errors.New("coupon not found")
	}

	if !s.isCouponApplicable(coupon, cart) {
		return nil, errors.New("coupon is not applicable")
	}

	strategy := s.strategyFactory.GetStrategy(coupon.Type)
	if strategy == nil {
		return nil, errors.New("unsupported coupon type")
	}

	updatedCart, err := strategy.ApplyCoupon(coupon, cart)
	if err != nil {
		return nil, err
	}

	// Increment usage count if there is a usage limit
	if coupon.UsageLimit > 0 {
		err := s.repo.IncrementUsageCount(coupon.ID)
		if err != nil {
			return nil, err
		}
	}

	return updatedCart, nil
}

func (s *couponService) isCouponApplicable(coupon *models.Coupon, cart *models.Cart) bool {
	// Check expiration date
	if coupon.ExpirationDate != nil && time.Now().After(*coupon.ExpirationDate) {
		return false
	}

	// Check usage limit
	if coupon.UsageLimit > 0 && coupon.UsedCount >= coupon.UsageLimit {
		return false
	}

	// Check user-specific coupon
	if coupon.Type == models.UserSpecific && !s.isCouponForUser(coupon, cart.UserID) {
		return false
	}

	// Additional checks can be added here
	return true
}

func (s *couponService) isCouponForUser(coupon *models.Coupon, userID uint) bool {
	for _, id := range coupon.Users {
		if id == userID {
			return true
		}
	}
	return false
}
