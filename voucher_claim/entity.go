package voucherclaim

import (
	"time"

	"github.com/google/uuid"
)

type VoucherClaim struct {
	ID        int
	UserID    uuid.UUID
	VoucherID int
	CreatedAt time.Time
}
