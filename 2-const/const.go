package main

import "fmt"

const (
	b = 10
	c = 20
)

const (
	d = iota
	e
	f = 2 * iota
	g
)

func main() {
	const a int = 1
	fmt.Println("a = ", a)
	fmt.Println("b = ", b, "c = ", c)
	fmt.Println("d = ", d, "e = ", e)
	fmt.Println("f = ", f, "g = ", g)
}
