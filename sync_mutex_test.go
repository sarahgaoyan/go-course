package course

import (
	"fmt"
	"sync"
	"testing"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func TestSync(t *testing.T) {
	t.Run("test syc", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		go worker(&wg)
		go worker(&wg)
		wg.Wait()

		fmt.Println(total.value)
	})
}
