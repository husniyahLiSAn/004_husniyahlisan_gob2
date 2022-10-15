package main

import (
	"ass-02/config"
	"ass-02/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}
