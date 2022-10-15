package main

import (
	"ass-03/config"
	"ass-03/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}
