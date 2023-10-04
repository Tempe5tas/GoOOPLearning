package Timeout_Control

import (
	"fmt"
	"testing"
	"time"
)

func Service() chan string {
	ret := make(chan string)
	go func() {
		fmt.Println("Service 1 started.")
		//time.Sleep(time.Millisecond * 114514)
		time.Sleep(time.Millisecond * 114)
		ret <- "Cyka blyat"
		time.Sleep(time.Millisecond * 514)
		fmt.Println("Service 1 done.")
	}()
	return ret
}

func TestTimeout(t *testing.T) {
	select {
	case <-Service():
		t.Log("Well-functioning")
	case <-time.After(time.Millisecond * 1919):
		t.Log("Timed out")
	}
	time.Sleep(time.Millisecond * 810)
}
