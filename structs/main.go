package main

import "fmt"

type Address struct {
	StreetAdress string
	Number       int
	City         string
	State        string
}

type Customer struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	jorge := Customer{
		Name:   "Jorge",
		Age:    22,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t \n", jorge.Name, jorge.Age, jorge.Active)

	jorge.Active = false
	jorge.City = "Santo Angelo"

	fmt.Printf("Name: %s, Age: %d, Active: %t, City: %s \n", jorge.Name, jorge.Age, jorge.Active, jorge.City)
}
