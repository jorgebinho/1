package main

import "fmt"

func main() {
	fmt.Println(sum(50, 3, 70, 35, 78, 743))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
