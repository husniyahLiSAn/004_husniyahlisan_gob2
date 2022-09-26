package main

//import package
import (
	"fmt"
	"os"
	"strconv"
)

// struct person

type Person struct {
	// struct initialization
	Number  int
	Name    string
	Address string
	Job     string
	Reason  string
}

// function main
func main() {
	// input number to variable args
	var arguments = os.Args
	// convert string to integer
	n, _ := strconv.Atoi(arguments[1])

	// variable participants
	var participants = []Person{
		{Number: 1, Name: "Dadah Taufik", Address: "Solo", Job: "Fullstack developer", Reason: "Belajar dari pengalaman"},
		{Number: 2, Name: "Indra Permana", Address: "Yogyakarta", Job: "Fullstack developer", Reason: "Tertarik dari pengalaman dengan project Golang"},
		{Number: 3, Name: "Steve Watt", Address: "Sleman", Job: "Backend developer", Reason: "Karena ingin terus belajar"},
		{Number: 4, Name: "Caroline", Address: "Lampung", Job: "Fullstack developer", Reason: "Awalnya terpaksa belajar karena project, tapi saya disini karena ingin belajar fundamentalnya."},
		{Number: 5, Name: "Aditya", Address: "Bandung", Job: "Mahasiswa semester akhir", Reason: "Ingin menambah portofolio dan achievement yang dibutuhkan untuk bekerja"},
		{Number: 6, Name: "Annisa Dewi", Address: "Bogor", Job: "Mahasiswa semester akhir", Reason: "Karena banyak peluang yang terbuka dengan Golang"},
		{Number: 7, Name: "Khoirul Pratama", Address: "Nganjuk", Job: "Backend developer", Reason: "Ingin belajar lebih dalam lagi tentang Golang"},
	}

	// generateParticipants as function with partisipan & num as parameter
	generateParticipant(participants, n)
}

func generateParticipant(p []Person, n int) {
	switch {
	case n != 0 && n <= len(p):
		fmt.Println("Number :", p[n-1].Number)
		fmt.Println("Name :", p[n-1].Name)
		fmt.Println("Address :", p[n-1].Address)
		fmt.Println("Job :", p[n-1].Job)
		fmt.Println("Reason :", p[n-1].Reason)
	default:
		fmt.Println("There's no data with that person")
	}
}
