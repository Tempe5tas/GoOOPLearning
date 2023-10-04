package Any_Response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 114)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	num := 10
	// Incorrect case, cause goroutine leak.
	// Use buffered channel instead.
	// ch := make(chan string)
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Goroutines count: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Millisecond * 1919)
	t.Log("Goroutines count: ", runtime.NumGoroutine())
}
