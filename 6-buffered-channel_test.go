package learngogoroutines

import (
	"fmt"
	"testing"
)

func BufferChannel(channel chan string) {

	channel <- "Rasyidev Pro"
	channel <- "Kim Tae Ri"

	fmt.Println("Berhasil memasukkan data ke dalam channel")
}

func TestBufferChannel(t *testing.T) {
	ch := make(chan string, 2)
	defer close(ch)

	go BufferChannel(ch)
	fmt.Println("Channel Capacity:", cap(ch))
	fmt.Println("Channel Length:", len(ch))
	data := <-ch
	data2 := <-ch

	fmt.Println("Selesai", data, data2)
}

/*
$ go test -v -run TestBufferChannel
=== RUN   TestBufferChannel
Channel Capacity: 2
Channel Length: 2
Selesai Rasyidev Pro Kim Tae Ri
Berhasil memasukkan data ke dalam channel
--- PASS: TestBufferChannel (0.00s)
PASS
ok      learn-go-goroutines     0.037s
*/
