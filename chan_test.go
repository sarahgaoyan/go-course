package course

import (
	"fmt"
	"testing"
)

// no channel buffer
var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello world"
	done <- true
}

func TestChan(t *testing.T) {
	t.Run("test go routing with chan", func(t *testing.T) {
		go aGoroutine()
		<-done
		fmt.Println(msg)
	})
}
