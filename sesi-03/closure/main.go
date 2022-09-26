package main

import (
	"fmt"
	"strings"
)

func main() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{7, 95, 23, 48, 56, 12, 34}
	fmt.Println(evenNumbers(numbers...))

	space()

	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}

	fmt.Println(isPalindrome("katak"))

	space()

	var studentLists = []string{"Airfell", "Nanda", "Mailo", "Angela"}

	var findstud = findStudent(studentLists)

	fmt.Println(findstud("Angela"))

	space()

	var findnum = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", findnum)

	space()

}

func space() {
	fmt.Println("==================================================================")
}

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position)
	}
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}
