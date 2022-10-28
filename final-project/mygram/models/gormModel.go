package models

import "time"

type GormModel struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp;column:updated_at" json:"updated_at,omitempty"`
}
