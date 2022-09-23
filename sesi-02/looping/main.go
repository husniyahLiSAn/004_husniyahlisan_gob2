package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	array := []int{7, 2, 6, 4, 9}
	var result []int
	fruits := map[string]string{"1": "apple", "2": "banana"}

	// Looping for loop arguments
	for i := len(array) - 1; i >= 0; i-- {
		result = append(result, array[i])
	}
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Reversed array: %v\n", result)

	// Looping for in range arguments
	fmt.Println("Fruits:")
	for k, v := range fruits {
		l := utf8.RuneCountInString(v)
		fmt.Printf("%s -> %s -> length: %d\n", k, v, l)
	}
}
