package learngogoroutines

import (
	"fmt"
	"testing"
)

func ManyGoroutines(num int) {
	fmt.Printf("Goroutines no-%v\n", num)
}

func TestManyNoGoroutines(t *testing.T) {
	totalGoroutine := 1000000 // 1 juta goroutines
	for i := 0; i < totalGoroutine; i++ {
		ManyGoroutines(i)
	}
}

func TestManyGoroutines(t *testing.T) {
	totalGoroutine := 1000000 // 1 juta goroutines
	for i := 0; i < totalGoroutine; i++ {
		go ManyGoroutines(i)
	}
}

/*
$ go test -v -run TestManyNoGoroutines
Goroutines no-1
Goroutines no-2
Goroutines no-3
.....
Goroutines no-999998
Goroutines no-999999
--- PASS: TestManyNoGoroutines (164.14s)
PASS
ok      learn-go-goroutines     164.185s
*/

/*
$ go test -v -run TestManyNoGoroutines
Goroutines no-3
Goroutines no-1
Goroutines no-2
.....
Goroutines no-999994
Goroutines no-999996
Goroutines no-999998
Goroutines no-999997
Goroutines no-999999
Goroutines no-35213
*/
