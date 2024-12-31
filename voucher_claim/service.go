package voucherclaim

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateVoucherClaim(userID uuid.UUID, voucherID int) (VoucherClaim, error)
	GetVoucherClaimByID(id int) (VoucherClaim, error)
	GetVoucherClaimByUserID(id uuid.UUID) ([]VoucherClaim, error)
	GetAllVoucherClaims() ([]VoucherClaim, error)
	UpdateVoucherClaim(id int, userID uuid.UUID, voucherID int) (VoucherClaim, error)
	DeleteVoucherClaim(id int) error
}

type service struct {
	repository VoucherClaimRepository
}

func NewService(repository VoucherClaimRepository) *service {
	return &service{repository}
}

func (s *service) CreateVoucherClaim(userID uuid.UUID, voucherID int) (VoucherClaim, error) {
	voucherClaim := VoucherClaim{
		UserID:    userID,
		VoucherID: voucherID,
	}
	return s.repository.Create(voucherClaim)
}

func (s *service) GetVoucherClaimByID(id int) (VoucherClaim, error) {
	return s.repository.FindByID(id)
}
func (s *service) GetVoucherClaimByUserID(userID uuid.UUID) ([]VoucherClaim, error) {
	return s.repository.FindByUserID(userID)
}

func (s *service) GetAllVoucherClaims() ([]VoucherClaim, error) {
	return s.repository.FindAll()
}

func (s *service) UpdateVoucherClaim(id int, userID uuid.UUID, voucherID int) (VoucherClaim, error) {
	voucherClaim, err := s.repository.FindByID(id)
	if err != nil {
		return VoucherClaim{}, errors.New("voucher claim not found")
	}

	voucherClaim.UserID = userID
	voucherClaim.VoucherID = voucherID
	return s.repository.Update(voucherClaim)
}

func (s *service) DeleteVoucherClaim(id int) error {
	return s.repository.Delete(id)
}
