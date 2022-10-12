package database

import (
	"fmt"
	"log"
	"web-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "puspakirana"
	password = "MyPostgre@123"
	dbPort   = "5432"
	dbName   = "db_go_rest"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbName)

	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Product{})

}

func GetDB() *gorm.DB {
	return db
}
