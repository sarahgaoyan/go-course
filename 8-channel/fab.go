package main

import (
	"fmt"
	"time"
)

var c = make(chan int)
var q = make(chan int)

func main() {

	go func() {
		defer fmt.Println("go routine end")
		for i := 0; i < 3; i++ {
			c <- i
			//fmt.Println(c <-i)
		}
		q <- 0
	}()

	fab()
}

func fab() {
	x, y := 1, 1
	for {
		select {
		case <-c:
			tmp := x
			x = y
			y = tmp + y
			fmt.Println("x: ", x, "y: ", y)
		case <-q:
			time.Sleep(1 * time.Second)
			return
		default:
		}
	}
}
