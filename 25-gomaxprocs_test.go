package learngogoroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU\t:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread\t:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine\t:", totalGoroutine)
}

/*
$ go test -v -run TestGetGomaxprocs
=== RUN   TestGetGomaxprocs
Total CPU       : 8
Total Thread    : 8
Total Goroutine : 2
--- PASS: TestGetGomaxprocs (0.00s)
PASS
ok      learn-go-goroutines     0.053s
*/
