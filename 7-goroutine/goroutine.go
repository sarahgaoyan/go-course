package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("defer A")
		func() {
			defer fmt.Println("defer B")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
