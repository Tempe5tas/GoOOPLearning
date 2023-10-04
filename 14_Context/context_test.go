package Context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // Received any signal, or channel closed (return default value).
		return true
	default: // Nothing to receive.
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	// Create five goroutines, waiting for channel signal to abort.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			for {
				if Cancelled(ctx) {
					break
				} else {
					time.Sleep(time.Millisecond * 114)
				}
			}
			fmt.Println("Received stop signal")
			wg.Done()
		}(i, ctx)
	}
	cancel()
	wg.Wait()
}
