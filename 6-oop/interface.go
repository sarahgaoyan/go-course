package main

import "fmt"

func printInterface(a interface{}) {
	fmt.Println(a)

	_, ok := a.(string)
	if !ok {
		fmt.Println("a is not string")
	} else {
		fmt.Println("a is string")
	}
}

func main() {
	a := "abc"
	printInterface(a)

	b := 234
	printInterface(b)

	c := 3.14
	printInterface(c)

	d := true
	printInterface(d)
}
