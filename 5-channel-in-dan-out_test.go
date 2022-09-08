package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(3 * time.Second)
	channel <- "Rasyidev Pro"
	fmt.Println("Berhasil memasukkan data ke dalam channel")
}

func OnlyOut(channel <-chan string) {
	time.Sleep(3 * time.Second)
	data := <-channel
	fmt.Println("Berhasil menerima data dari channel:", data)
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("Running TestInOutChannel function")
}

/*
$ go test -v -run TestInOutChannel
=== RUN   TestInOutChannel
Berhasil memasukkan data ke dalam channel
Berhasil menerima data dari channel: Rasyidev Pro
Running TestInOutChannel function
--- PASS: TestInOutChannel (5.00s)
PASS
ok      learn-go-goroutines     5.045s
*/
