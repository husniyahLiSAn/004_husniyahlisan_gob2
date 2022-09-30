// SLide Reflect
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Type variable :", reflectValue.Type())
	fmt.Println("Value variable :", reflectValue.Interface())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Value variable :", reflectValue.Int())
	}

	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)
	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value{reflect.ValueOf("wick")})

	fmt.Println("Name :", s1.Name)
}
