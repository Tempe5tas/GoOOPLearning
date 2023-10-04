package Singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// Lazy Singleton

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingleObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
