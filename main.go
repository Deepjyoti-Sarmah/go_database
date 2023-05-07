package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	PinCode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	// fmt.Println("Go Database")
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "233384", "Myrl Tech", Address{"Silchar", "Assam", "India", "788003"}},
		{"Deep", "22", "843445", "World Tech", Address{"Guwahati", "Assam", "India", "781023"}},
		{"Akhil", "33", "238433", "Akhil Tech", Address{"Bengaloru", "Kernataka", "India", "410015"}},
		{"Ania", "25", "788992", "Myrl Tech", Address{"Silchar", "Assam", "India", "788003"}},
		{"Jeep", "25", "355684", "World Tech", Address{"Guwahati", "Assam", "India", "781023"}},
		{"Rohit", "29", "912684", "Akhil Tech", Address{"Bengaloru", "Kernataka", "India", "410015"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}

	fmt.Println((allusers))

	// if err := db.Delete("user", "john"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
