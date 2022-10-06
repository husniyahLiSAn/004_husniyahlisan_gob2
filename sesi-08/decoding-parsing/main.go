package main

import (
	"encoding/json"
	"fmt"
)

// Decoding JSON To Struct
// Decoding JSON Array To Slice Of Struct
type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	// Decoding JSON To Struct
	var jsonString = `
	{
		"full_name" : "Husniyah Lisan",
		"email": "husniyah@gmail.com",
		"age": 22
	}
	`

	// json.Unmarshal untuk mendecode data JSON kepada struct Employee
	var result Employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("=====================================")
	fmt.Println("Decoding JSON To Struct")
	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
	fmt.Println("=====================================")

	// Decoding JSON To Map
	var jsonString2 = `
	{
		"full_name" : "Husniyah Lisan",
		"email": "husniyah@gmail.com",
		"age": 22
	}
	`
	var result2 map[string]interface{}

	var err2 = json.Unmarshal([]byte(jsonString2), &result2)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Decoding JSON To Map")
	fmt.Println("full_name:", result2["full_name"])
	fmt.Println("email:", result2["email"])
	fmt.Println("age:", result2["age"])
	fmt.Println("=====================================")

	// Decoding JSON To Empty Interface
	var jsonString3 = `
	{
		"full_name" : "Husniyah Lisan",
		"email": "husniyah@gmail.com",
		"age": 22
	}
	`

	var temp interface{}

	var err3 = json.Unmarshal([]byte(jsonString3), &temp)
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}
	var result3 = temp.(map[string]interface{})

	fmt.Println("Decoding JSON To Empty Interface")
	fmt.Println("full_name:", result3["full_name"])
	fmt.Println("email:", result3["email"])
	fmt.Println("age:", result3["age"])
	fmt.Println("=====================================")

	// Decoding JSON Array To Slice Of Struct
	var jsonString4 = `[
			{
				"full_name" : "Husniyah Lisan",
				"email": "husniyah@gmail.com",
				"age": 22
			},
			{
				"full_name" : "Amelia",
				"email": "amelia@gmail.com",
				"age": 23
			}
		]
	`

	var result4 []Employee

	var err4 = json.Unmarshal([]byte(jsonString4), &result4)
	if err4 != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result4 {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
	fmt.Println("=====================================")
}
