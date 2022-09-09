package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}
	group := sync.WaitGroup{}

	pool.Put("Rasyidev Pro")
	pool.Put("Kim Taeri")
	pool.Put("Bae Suzy")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			res := pool.Get()
			// res2 := pool.Get()
			// res3 := pool.Get()

			fmt.Println(res)
			// fmt.Println(res2)
			// fmt.Println(res3)

			pool.Put(res)
			// pool.Put(res2)
			// pool.Put(res3)
			group.Done()
		}()
	}
	group.Wait()
}

/*
$ go test -v -run TestPool
=== RUN   TestPool
Rasyidev Pro
Kim Taeri
Bae Suzy
<nil>
<nil>
<nil>
<nil>
<nil>
<nil>
<nil>
--- PASS: TestPool (0.00s)
PASS
ok      learn-go-goroutines     0.038s
*/
