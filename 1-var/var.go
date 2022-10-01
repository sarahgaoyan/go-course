package main

import "fmt"

var e int = 4

//f := 5

var g, h = 6, "six"

var (
	i int    = 7
	j string = "seven"
)

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 1
	var c = 2
	fmt.Println("b = ", b, "c = ", c)

	d := 3
	fmt.Printf("d = %d, type of d = %T\n", d, d)

	fmt.Printf("e = %d\n", e)
	//fmt.Printf("f = %d\n", f)  // : definition not allow for global variables

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "j = ", j)
}
