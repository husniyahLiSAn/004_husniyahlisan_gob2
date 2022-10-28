package models

import (
	"errors"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email    string `gorm:"column:email;type:varchar(150);not null;uniqueIndex" json:"email" form:"email" valid:"email,required"`
	Username string `gorm:"column:username;type:varchar(150);not null; unique" json:"username" form:"username" valid:"required"`
	Password string `gorm:"column:password;type:varchar(150);not null" json:"password" form:"password" valid:"required, minstringlength(6)"`
	Age      int64  `gorm:"column:age;type:int;not null" json:"age" form:"age" valid:"required,numeric"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	errAge := ValidateAge(u)
	if errAge != nil {
		err = errAge
		return
	}
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func ValidateAge(user *User) error {
	if user.Age < 8 {
		return errors.New("age: minimal age is 8")
	}
	return nil
}
