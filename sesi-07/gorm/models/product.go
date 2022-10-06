package models

import (
	"time"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique;type:varchar(191)"`
	Brand     string `gorm:"not null;unique;type:varchar(191)"`
	UserID    uint   // Foreign Key ID from User
	CreatedAt time.Time
	UpdateAt  time.Time
}
