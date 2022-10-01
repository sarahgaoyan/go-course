package course

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncOnce(t *testing.T) {
	t.Run("test syc once", func(t *testing.T) {
		o := &sync.Once{}
		for i := 0; i < 10; i++ {
			o.Do(func() {
				fmt.Println("only once")
			})
		}
	})
}
