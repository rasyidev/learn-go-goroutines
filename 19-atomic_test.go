package learngogoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", x)
}

/*
$ go test -v -run TestAtomic
=== RUN   TestAtomic
Counter: 100000
--- PASS: TestAtomic (0.00s)
PASS
ok      learn-go-goroutines     0.061s
*/
