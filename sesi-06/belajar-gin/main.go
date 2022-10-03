package main

import "belajar-gin/router"

func main() {
	var PORT = ":8082"

	router.StartServer().Run(PORT)
}
