package learngogoroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

var data sync.Map

func addToMap(v int) {
	data.Store(strconv.Itoa(v), v*10)
}

func TestMap(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go addToMap(i)
		group.Done()
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

}

/*
90: 900
7: 70
.....
4: 40
45: 450
--- PASS: TestMap (0.02s)
PASS
ok      learn-go-goroutines     0.067s
*/
