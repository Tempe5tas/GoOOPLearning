package Task_Cancel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Cancelled(signal chan bool) bool {
	select {
	case <-signal: // Received any signal, or channel closed (return default value).
		return true
	default: // Nothing to receive.
		return false
	}
}

// Only cancel one goroutine.
func Cancel1(signal chan bool) {
	signal <- true // Send signal to channel
}

// Cancel all goroutine.
func Cancel2(signal chan bool) {
	close(signal) // Close channel
}

func TestCancel(t *testing.T) {
	signal := make(chan bool)
	wg := sync.WaitGroup{}

	// Create five goroutines, waiting for channel signal to abort.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, signal chan bool) {
			for {
				if Cancelled(signal) {
					break
				} else {
					time.Sleep(time.Millisecond * 114)
				}
			}
			fmt.Println("Received stop signal")
			wg.Done()
		}(i, signal)
	}

	//Cancel1(signal)
	Cancel2(signal)
	wg.Wait()
}
