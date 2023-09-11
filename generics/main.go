package main

type Numero interface {
	~int | float64
}

func Soma[T Numero](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{
		"Jorge":   1000,
		"Bruno":   2000,
		"Claudio": 3000}

	println(Soma(m))
}
