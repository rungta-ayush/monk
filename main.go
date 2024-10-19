package main

import (
	"log"

	"coupon-api/service/strategies"

	"coupon-api/handlers"
	"coupon-api/repositories"
	"coupon-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the repository
	couponRepo, err := repositories.NewCouponRepository("data/coupons.json")
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize the strategy factory
	strategyFactory := strategies.NewCouponStrategyFactory()

	// Initialize the service
	couponService := services.NewCouponService(couponRepo, strategyFactory)

	// Initialize the handler
	couponHandler := handlers.NewCouponHandler(couponService)

	// Set up the router
	router := gin.Default()

	// Define the routes
	router.POST("/coupons", couponHandler.CreateCoupon)
	router.GET("/coupons", couponHandler.GetCoupons)
	router.GET("/coupons/:id", couponHandler.GetCouponByID)
	router.PUT("/coupons/:id", couponHandler.UpdateCoupon)
	router.DELETE("/coupons/:id", couponHandler.DeleteCoupon)
	router.POST("/applicable-coupons", couponHandler.GetApplicableCoupons)
	router.POST("/apply-coupon/:id", couponHandler.ApplyCoupon)
	// Serve the swagger.yaml file
	router.Static("/docs", "./docs")

	// Serve the Swagger UI files
	router.Static("/swagger", "./swagger-ui")
	// Start the server
	router.Run(":8080")
}
