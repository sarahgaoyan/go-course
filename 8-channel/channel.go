package main

import "fmt"

func main() {
	//c := make(chan int)
	c := make(chan int, 2)

	go func() {
		defer fmt.Println("go routine end")
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("go routine send", i)
		}
		close(c)
	}()

	//for i := 0; i < 2; i++ {
	//	num := <-c
	//	fmt.Println("main routine num = ", num)
	//}

	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println("main routine num = ", data)
	//	} else {
	//		break
	//	}
	//}

	for data := range c {
		fmt.Println("main routine num = ", data)
	}

	fmt.Println("main routine...")
}
