package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	jorge := Customer{
		Name:   "Jorge",
		Age:    22,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t \n", jorge.Name, jorge.Age, jorge.Active)

	jorge.Active = false

	fmt.Printf("Name: %s, Age: %d, Active: %t \n", jorge.Name, jorge.Age, jorge.Active)
}
