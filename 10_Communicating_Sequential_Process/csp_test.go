package Communicating_Sequential_Process

import (
	"fmt"
	"testing"
	"time"
)

func Service1() {
	fmt.Println("Service 1 started.")
	time.Sleep(time.Millisecond * 114)
	fmt.Println("Service 1 done.")
}

func Service2() {
	fmt.Println("Service 2 started.")
	time.Sleep(time.Millisecond * 1919)
	fmt.Println("Service 2 done.")
}

func AsyncService() chan string {
	//Non buffer
	//ret := make(chan string)

	//buffered channel, capacity=1
	ret := make(chan string, 1)
	go func() {
		fmt.Println("Async service started.")
		ret <- "Cyka"
		//time.Sleep(time.Millisecond * 114)
		fmt.Println("Async service done.")
	}()
	return ret
}

func TestService(t *testing.T) {
	Service1()
	Service2()
}

func TestAsyncService(t *testing.T) {
	ret := AsyncService()
	Service2()
	fmt.Println(<-ret)
	time.Sleep(time.Millisecond * 114)
}
