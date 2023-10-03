package Shared_Memory

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSharedMemory(t *testing.T) {
	var num int
	for i := 0; i < 114514; i++ {
		go func() {
			num++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}

func TestMutex(t *testing.T) {
	var num int
	var lock sync.Mutex
	for i := 0; i < 114514; i++ {
		go func() {
			defer func() {
				lock.Unlock()
			}()
			lock.Lock()
			num++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}

func TestWaitGroup(t *testing.T) {
	var num int
	var lock sync.Mutex
	var group sync.WaitGroup
	for i := 0; i < 114514; i++ {
		group.Add(1)
		go func() {
			defer func() {
				lock.Unlock()
				group.Done()
			}()
			lock.Lock()
			num++
		}()
	}
	group.Wait()
	fmt.Println(num)
}
