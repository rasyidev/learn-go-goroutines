package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Berhasil dieksekusi")
		group.Done()
	})

	group.Wait()
}

/*
=== RUN   TestAfterFunc
Berhasil dieksekusi
--- PASS: TestAfterFunc (3.01s)
PASS
ok      learn-go-goroutines     3.055s
*/
