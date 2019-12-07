package sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new objc")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC是会清除掉synv.pool里面的缓存对象
	v1 := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new objc")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
