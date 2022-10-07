package main

import "fmt"

func main() {
	a := 10
	b := 20

	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)
}

func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}
