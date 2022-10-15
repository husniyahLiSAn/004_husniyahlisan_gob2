package main

import (
	"web-app/database"
	"web-app/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}
