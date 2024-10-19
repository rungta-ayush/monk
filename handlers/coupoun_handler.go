package handlers

import (
	"net/http"
	"strconv"

	"coupon-api/models"
	"coupon-api/services"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	service services.CouponService
}

func NewCouponHandler(service services.CouponService) *CouponHandler {
	return &CouponHandler{service: service}
}

func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateCoupon(&coupon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, coupon)
}

func (h *CouponHandler) GetCoupons(c *gin.Context) {
	coupons, err := h.service.GetAllCoupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupons)
}

func (h *CouponHandler) GetCouponByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon ID"})
		return
	}
	coupon, err := h.service.GetCouponByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "coupon not found"})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) UpdateCoupon(c *gin.Context) {
	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon ID"})
		return
	}
	coupon.ID = uint(id)
	if err := h.service.UpdateCoupon(&coupon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) DeleteCoupon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon ID"})
		return
	}
	if err := h.service.DeleteCoupon(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon deleted"})
}

func (h *CouponHandler) GetApplicableCoupons(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	applicableCoupons, err := h.service.GetApplicableCoupons(&cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applicable_coupons": applicableCoupons})
}

func (h *CouponHandler) ApplyCoupon(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon ID"})
		return
	}
	updatedCart, err := h.service.ApplyCoupon(uint(id), &cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated_cart": updatedCart})
}
