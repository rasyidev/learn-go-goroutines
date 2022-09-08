package learngogoroutines

import (
	"fmt"
	"testing"
)

func RangeChannel(channel chan string) {
	buah := []string{
		"Apel", "Mangga", "Semangka", "Markisa", "Srikaya", "Pisang",
	}

	// channel bisa diisi banyak data karena func RangeChannel dipanggil menggunakan goroutine
	for _, b := range buah {
		channel <- b
	}
	close(channel)
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)
	go RangeChannel(ch)

	for buah := range ch {
		fmt.Println("Menerima data:", buah)
	}
}

/*
$ go test -v -run TestRangeChannel
=== RUN   TestRangeChannel
Menerima data: Apel
Menerima data: Mangga
Menerima data: Semangka
Menerima data: Markisa
Menerima data: Srikaya
Menerima data: Pisang
--- PASS: TestRangeChannel (0.00s)
PASS
ok      learn-go-goroutines     0.038s
*/
