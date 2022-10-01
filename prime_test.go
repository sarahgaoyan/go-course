package course

import (
	"context"
	"fmt"
	"testing"
)

func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func TestPrime(t *testing.T) {
	t.Run("test prime goroutine", func(t *testing.T) {
		// control background go routine by context
		ctx, cancel := context.WithCancel(context.Background())
		ch := GenerateNatural(ctx)
		for i := 0; i < 100; i++ {
			prime := <-ch
			fmt.Printf("%v: %v\n\n", i+1, prime)
			ch = PrimeFilter(ctx, ch, prime)
		}
		cancel()
	})
}
