package Sync_Pool

import (
	"sync"
	"testing"
)

// Warning:
// sync.Pool lifetime is affected by GC, cache will be flushed during GC.

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("New object created.")
			return 114514
		},
	}
	v := pool.Get().(int)
	t.Log(v)
	pool.Put("Cyka blyat")
	//runtime.GC()
	if v1, ok := pool.Get().(string); !ok {
		t.Error("mismatched types")
	} else {
		t.Log(v1)
	}

}

func TestSycPoolGoroutines(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("New object created.")
			return 114514
		},
	}
	pool.Put(114)
	pool.Put(514)
	pool.Put(1919)
	pool.Put(810)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
		wg.Wait()
	}
}
