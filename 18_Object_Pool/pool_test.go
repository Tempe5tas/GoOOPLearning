package Object_Pool

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

type ResuableObj struct {
	cyka  int
	blyat string
}

type ObjPool struct {
	bufChan chan *ResuableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResuableObj, num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &ResuableObj{}
	}
	return &objPool
}

func (pool *ObjPool) GetObj(timeout time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-pool.bufChan:
		return ret, nil
	case <-time.After(timeout): //timeout control
		return nil, errors.New("timed out")
	}
}

func (pool *ObjPool) ReleaseObj(obj *ResuableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflowed")
	}
}

func TestObjPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Millisecond * 114); err != nil {
			t.Error(err)
		} else {
			v.cyka = rand.Intn(10)
			v.blyat = string(rune(rand.Intn(114514)))
			t.Log(v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("All done.")
}
