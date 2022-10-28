package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"column:title;type:varchar(255);not null;" json:"title" form:"title" valid:"required"`
	Caption  string `gorm:"column:caption;type:text;not null;" json:"caption" form:"caption" valid:"required"`
	PhotoUrl string `gorm:"column:photo_url;type:text;not null;" json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint64
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
