package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"coupon-api/models"
)

type CouponRepository interface {
	CreateCoupon(coupon *models.Coupon) error
	GetAllCoupons() ([]models.Coupon, error)
	GetCouponByID(id uint) (*models.Coupon, error)
	UpdateCoupon(coupon *models.Coupon) error
	DeleteCoupon(id uint) error
	IncrementUsageCount(id uint) error
}

type couponRepository struct {
	filePath string
	coupons  []models.Coupon
	mutex    sync.Mutex
}

func NewCouponRepository(filePath string) (CouponRepository, error) {
	repo := &couponRepository{filePath: filePath}
	err := repo.loadCoupons()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *couponRepository) loadCoupons() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			r.coupons = []models.Coupon{}
			return nil
		}
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &r.coupons)
}

func (r *couponRepository) saveCoupons() error {
	data, err := json.MarshalIndent(r.coupons, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.filePath, data, 0644)
}

func (r *couponRepository) CreateCoupon(coupon *models.Coupon) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	coupon.ID = uint(len(r.coupons) + 1)
	r.coupons = append(r.coupons, *coupon)
	return r.saveCoupons()
}

func (r *couponRepository) GetAllCoupons() ([]models.Coupon, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.coupons, nil
}

func (r *couponRepository) GetCouponByID(id uint) (*models.Coupon, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, coupon := range r.coupons {
		if coupon.ID == id {
			return &coupon, nil
		}
	}
	return nil, errors.New("coupon not found")
}

func (r *couponRepository) UpdateCoupon(coupon *models.Coupon) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, c := range r.coupons {
		if c.ID == coupon.ID {
			r.coupons[i] = *coupon
			return r.saveCoupons()
		}
	}
	return errors.New("coupon not found")
}

func (r *couponRepository) DeleteCoupon(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, c := range r.coupons {
		if c.ID == id {
			r.coupons = append(r.coupons[:i], r.coupons[i+1:]...)
			return r.saveCoupons()
		}
	}
	return errors.New("coupon not found")
}

func (r *couponRepository) IncrementUsageCount(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, c := range r.coupons {
		if c.ID == id {
			r.coupons[i].UsedCount++
			return r.saveCoupons()
		}
	}
	return errors.New("coupon not found")
}
