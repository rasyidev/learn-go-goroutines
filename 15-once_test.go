package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
	fmt.Println("Counter berjalan 1x")
}

func TestOnce(t *testing.T) {
	group := sync.WaitGroup{}
	once := sync.Once{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			// kalau kita jalankan tanpa once.Do(), akan terjadi race condition
			// OnlyOnce()
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}

/*
$ go test -v -run TestOnce
=== RUN   TestOnce
Counter berjalan 1x
Counter: 1
--- PASS: TestOnce (0.00s)
PASS
ok      learn-go-goroutines     0.048s
*/
