package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	// kirim event setiap 1 detik, datanya berupa waktu pengiriman event
	ticker := time.NewTicker(1 * time.Second)

	// Stop ticker setelah 5 detik
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

/*
$ go test -v -run TestTicker
=== RUN   TestTicker
2022-09-10 10:06:38.4515396 +0700 +07 m=+1.021482201
2022-09-10 10:06:39.4425913 +0700 +07 m=+2.012533901
2022-09-10 10:06:40.4514777 +0700 +07 m=+3.021420301
2022-09-10 10:06:41.4444197 +0700 +07 m=+4.014362301
2022-09-10 10:06:42.4411448 +0700 +07 m=+5.011087401
exit status 0xc000013a
FAIL    learn-go-goroutines     14.120s
*/
