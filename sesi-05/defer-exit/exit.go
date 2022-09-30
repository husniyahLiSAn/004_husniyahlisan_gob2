package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
