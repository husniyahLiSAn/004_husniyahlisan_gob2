package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Starting")

	greet1("Angela", 8)
	space()
	greet2("John", "Street Lorem Ipsum")
	space()
	fmt.Println(greet3("Nice to know you here", []string{"Angela", "John", "Mike"}))
	space()
	profile("Angela", "Pasta", "Padang")
	space()
	fmt.Println(print("Angela", "Micella", "John", "Mike"))
	space()
	fmt.Println(calculate(21))
	space()
	fmt.Println(calculate2(54))
	space()
	fmt.Println(sum(5, 6, 1, 23, 9))
	space()
}

func greet1(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func space() {
	fmt.Println("==================================================================")
}

func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func calculate2(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
