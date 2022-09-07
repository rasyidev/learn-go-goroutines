package learngogoroutines

import (
	"fmt"
	"testing"
)

func ManyGoroutines(num int) {
	fmt.Printf("Goroutines no-%v\n", num)
}

func TestManyNoGoroutines(t *testing.T) {
	totalGoroutine := 100000 // seratus ribu goroutines
	for i := 0; i < totalGoroutine; i++ {
		ManyGoroutines(i)
	}
}

func TestManyGoroutines(t *testing.T) {
	totalGoroutine := 100000 // seratus ribu goroutines
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
Goroutines no-99997
Goroutines no-99998
Goroutines no-99999
--- PASS: TestManyNoGoroutines (14.10s)
PASS
ok      learn-go-goroutines     14.143s
*/

/*
$ go test -v -run TestManyNoGoroutines
Goroutines no-3
Goroutines no-1
Goroutines no-2
.....
Goroutines no-99995
Goroutines no-99997
Goroutines no-99998
Goroutines no-99999
--- PASS: TestManyGoroutines (1.10s)
PASS
ok      learn-go-goroutines     15.334s
*/
