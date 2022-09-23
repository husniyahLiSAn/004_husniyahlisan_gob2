package main

import "fmt"

func main() {
	// Conditional arguments

	// If...else...
	var currentYear = 2021
	if age := currentYear - 2008; age < 17 {
		fmt.Printf("Kamu belum boleh membuat kartu sim dan umurmu saat ini %d\n", age)
	} else {
		fmt.Printf("Kamu sudah boleh membuat kartu sim dan umurmu saat ini %d\n", age)
	}

	// If...else nested...
	nilai := 78
	if nilai < 75 {
		fmt.Println("Kamu harus ikut kelas tambahan dan remedial")
	} else if nilai >= 75 && nilai <= 80 {
		fmt.Println("Nilai kamu sudah cukup bagus. Tetaplah rajin belajar")
	} else if nilai > 80 {
		fmt.Println("Nilai kamu sudah sangat bagus. Pertahankan")
	}

	// Switch case arguments
	switch {
	case nilai < 75:
		fmt.Println("C")
	case nilai >= 75 && nilai <= 80:
		fmt.Println("B")
	default:
		fmt.Println("A")
	}
}
