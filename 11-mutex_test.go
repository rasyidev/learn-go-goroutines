package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("X saat ini:", x)
}

/*
$ go test -v -run TestMutex
=== RUN   TestMutex
X saat ini: 10000
--- PASS: TestMutex (5.01s)
PASS
ok      learn-go-goroutines     5.052s
*/
