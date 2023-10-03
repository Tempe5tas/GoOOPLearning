package Goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//Data race occured.
		//go func() {
		//	fmt.Println(i)
		//}()

		//Function values pass by value,no data race.
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
