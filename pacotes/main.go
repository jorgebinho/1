package main

import (
	"fmt"
	"matematica/math"

	"github.com/google/uuid"
)

func main() {
	s := math.Soma(10, 20)
	fmt.Println("Resultado: ", s)
	fmt.Println(uuid.New())
}
