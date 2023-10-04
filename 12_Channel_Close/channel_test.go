package Channel_Close

import (
	"fmt"
	"sync"
	"testing"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//panic: send on closed channel
		//ch <- 114514
		wg.Done()
	}()
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

//Read on closed channel, return default value (0 for int).
//
//func BadConsumer(ch chan int, wg *sync.WaitGroup) {
//	go func() {
//		for i := 0; i < 114; i++ {
//			data := <-ch
//			fmt.Println(data)
//		}
//		wg.Done()
//	}()
//}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	Producer(ch, &wg)
	wg.Add(1)
	Consumer(ch, &wg)
	//wg.Add(1)
	//BadConsumer(ch, &wg)
	wg.Wait()
}
