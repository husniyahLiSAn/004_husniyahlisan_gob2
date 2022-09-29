package main

import (
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer cathErr()

	var username string

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validateLogin(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validateLogin(username string, password []byte) (string, error) {
	ul := len(username)
	pl := len(password)

	if username != "hlisa" && string(password) != "qw3rty00" {
		return "", fmt.Errorf("Invalid username/password!")
	} else if pl < 8 {
		return "", fmt.Errorf("Password must be at least 8 characters long")
	} else if ul < 5 {
		return "", fmt.Errorf("Password must be at least 5 characters long")
	}
	return "Login successful", nil
}

func cathErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running pervectly")
	}
}
