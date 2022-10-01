package design_patten

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestProCon(t *testing.T) {
	t.Run("test producer and consumer", func(t *testing.T) {
		ch := make(chan int, 64)
		go Producer(3, ch)
		go Producer(5, ch)
		go Consumer(ch)

		//time.Sleep(1 * time.Second)
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		fmt.Printf("quit (%v)\n", <-sig)
	})

}
