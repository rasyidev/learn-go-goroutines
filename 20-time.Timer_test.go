package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

/*
$ go test -v -run TestTimer
=== RUN   TestTimer
2022-09-10 09:43:56.2762499 +0700 +07 m=+0.006965901
2022-09-10 09:44:01.2868734 +0700 +07 m=+5.017589401
--- PASS: TestTimer (5.01s)
PASS
ok      learn-go-goroutines     5.055s
*/
