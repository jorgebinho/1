package main

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Jorge"
	e float64 = 1.2
	g ID      = 3
)

func main() {
	println(a)
	println(g)
}
