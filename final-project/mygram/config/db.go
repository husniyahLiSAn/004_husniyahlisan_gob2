package config

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	var (
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbPort   = os.Getenv("PGPORT")
		dbname   = os.Getenv("PGDATABASE")
	)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
