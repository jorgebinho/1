package main

import "fmt"

func main() {
	salary := map[string]int{"Jorge": 1700, "Rafael": 1500}
	fmt.Println(salary["Jorge"])
	delete(salary, "Rafael")
	salary["Nico"] = 5000
	fmt.Println(salary["Nico"])

	for name, salary := range salary {
		fmt.Printf("The salary of %s is of %d\n", name, salary)
	}
}
