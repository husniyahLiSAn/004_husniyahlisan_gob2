package main

import "fmt"

func main() {

	c := make(chan string)

	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	/*
		Hao, my name is Mailo
		Hao, my name is Airell
		Hao, my name is Nanda
		Keluaran mengikuti go introduce
	*/
	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hao, my name is %s", student)

	c <- result
}
