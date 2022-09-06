package learngogoroutines

import (
	"fmt"
	"testing"
)

func RunHelloRasyidev() {
	fmt.Println("Hello Rasyidev")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloRasyidev()
	fmt.Println("Selesai menjalankan function RunHelloRasyidev")
}

/*
$ go test -v
=== RUN   TestCreateGoroutine
Selesai menjalankan function RunHelloRasyidev
Hello Rasyidev
--- PASS: TestCreateGoroutine (0.00s)
PASS
ok      learn-go-goroutines     0.037s
*/
