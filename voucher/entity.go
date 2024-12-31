package voucher

import "time"

type Voucher struct {
	ID int
	Name string
	Foto string
	Category string
	Status bool
	CreatedAt time.Time
	UpdatedAt time.Time
}