package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);unique;not null"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
