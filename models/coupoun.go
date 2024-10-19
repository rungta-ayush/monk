package models

import "time"

type CouponType string

const (
	CartWise       CouponType = "cart-wise"
	ProductWise    CouponType = "product-wise"
	BxGy           CouponType = "bxgy"
	TimeBased      CouponType = "time-based"
	FirstTimeBuyer CouponType = "first-time-buyer"
	LimitedUse     CouponType = "limited-use"
	UserSpecific   CouponType = "user-specific"
	Referral       CouponType = "referral"
)

type Coupon struct {
	ID             uint        `json:"id"`
	Type           CouponType  `json:"type" binding:"required"`
	Details        interface{} `json:"details" binding:"required"`
	ExpirationDate *time.Time  `json:"expiration_date"`
	UsageLimit     uint        `json:"usage_limit,omitempty"`
	UsedCount      uint        `json:"used_count,omitempty"`
	Users          []uint      `json:"users,omitempty"` // User IDs for user-specific coupons
}
