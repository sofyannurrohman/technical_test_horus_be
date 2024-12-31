package voucher

import "gorm.io/gorm"

type Repository interface {
	Save(voucher Voucher) (Voucher, error)
	FindAll() ([]Voucher, error)
	FindByID(ID int) (Voucher, error)
	Update(voucher Voucher) (Voucher, error)
	Delete(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(voucher Voucher) (Voucher, error) {
	err := r.db.Create(&voucher).Error
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}
func (r *repository) FindAll() ([]Voucher, error) {
	var vouchers []Voucher
	err := r.db.Find(&vouchers).Error
	if err != nil {
		return vouchers, err
	}
	return vouchers, nil
}
func (r *repository) FindByID(ID int) (Voucher, error) {
	var voucher Voucher
	err := r.db.Where("id = ?", ID).Find(&voucher).Error
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}
func (r *repository) Update(voucher Voucher) (Voucher, error) {
	err := r.db.Save(&voucher).Error
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}

func (r *repository) Delete(ID int) error {
	err := r.db.Delete(&Voucher{}, ID).Error
	if err != nil {
		return err
	}
	return nil
}
