package config

import (
	"rest-api-project/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
