package voucherclaim

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherClaimRepository interface {
	Create(voucherClaim VoucherClaim) (VoucherClaim, error)
	FindByID(id int) (VoucherClaim, error)
	FindByUserID(id uuid.UUID) ([]VoucherClaim, error)
	FindAll() ([]VoucherClaim, error)
	Update(voucherClaim VoucherClaim) (VoucherClaim, error)
	Delete(id int) error
}

type voucherClaimRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) VoucherClaimRepository {
	return &voucherClaimRepository{db: db}
}

func (r *voucherClaimRepository) Create(voucherClaim VoucherClaim) (VoucherClaim, error) {
	err := r.db.Create(&voucherClaim).Error
	return voucherClaim, err
}

func (r *voucherClaimRepository) FindByID(id int) (VoucherClaim, error) {
	var voucherClaim VoucherClaim
	err := r.db.First(&voucherClaim, "id = ?", id).Error
	return voucherClaim, err
}
func (r *voucherClaimRepository) FindByUserID(userID uuid.UUID) ([]VoucherClaim, error) {
	var voucherClaim []VoucherClaim
	err := r.db.Find(&voucherClaim, "user_id = ?", userID).Error
	return voucherClaim, err
}

func (r *voucherClaimRepository) FindAll() ([]VoucherClaim, error) {
	var voucherClaims []VoucherClaim
	err := r.db.Find(&voucherClaims).Error
	return voucherClaims, err
}

func (r *voucherClaimRepository) Update(voucherClaim VoucherClaim) (VoucherClaim, error) {
	err := r.db.Save(&voucherClaim).Error
	return voucherClaim, err
}

func (r *voucherClaimRepository) Delete(id int) error {
	err := r.db.Delete(&VoucherClaim{}, "id = ?", id).Error
	return err
}
