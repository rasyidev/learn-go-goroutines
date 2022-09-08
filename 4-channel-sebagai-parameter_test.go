package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

/*
channel di parameter otomatis menggunakan pass by reference
tapi tidak perlu menambahkan *
*/
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rasyidev"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

/*
=== RUN   TestChannelAsParameter
Rasyidev
--- PASS: TestChannelAsParameter (7.03s)
PASS
ok      learn-go-goroutines     7.079s
*/
