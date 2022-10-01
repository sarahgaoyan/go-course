package course

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i < 10; i++ {
		atomic.AddUint64(&total2, i)
	}
}

func TestAtomic(t *testing.T) {
	t.Run("test atomic", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		go worker2(&wg)
		go worker2(&wg)
		wg.Wait()

		fmt.Println(total2)
	})
}
