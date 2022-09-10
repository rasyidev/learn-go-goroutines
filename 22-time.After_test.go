package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	// kirim data waktu setelah 5 detik
	channel := time.After(5 * time.Second)
	fmt.Println("Waktu saat ini:", time.Now())

	time := <-channel
	fmt.Println(time)
}

/*
=== RUN   TestAfter
Waktu saat ini: 2022-09-10 09:49:31.7081759 +0700 +07 m=+0.006549801
2022-09-10 09:49:36.7201405 +0700 +07 m=+5.018514401
--- PASS: TestAfter (5.01s)
PASS
ok      learn-go-goroutines     5.058s
*/
