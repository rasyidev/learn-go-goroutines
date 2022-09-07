package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	dataTobeSend := "Rasyidev"
	channel <- dataTobeSend
	receivedData := <-channel
	fmt.Println(<-channel)
	fmt.Println(receivedData)

	// close(channel) bisa close di sini atau pakai defer
}

// Ada pengirim tapi tidak ada penerima data
func TestChannelNoReceiver(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(3 * time.Second)
		channel <- "Rasyidev Pro"
		fmt.Println("Selesai mengirim data ke dalam channel")
	}()
}

/* Karena data dlm channel gak diambil2, udah keburu ditutup channelnya
$ go test -v -run TestChannelNoReceiver
=== RUN   TestChannelNoReceiver
panic: send on closed channel

goroutine 7 [running]:
learn-go-goroutines.TestChannelNoReceiver.func1()
      learn-go-goroutines/3-channel_test.go:26 +0x25
created by learn-go-goroutines.TestChannelNoReceiver
      learn-go-goroutines/3-channel_test.go:25 +0xa5
exit status 2
FAIL    learn-go-goroutines     0.053s
*/

// Ada penerima tapi tidak ada yang mengirim data ke dalam channel
func TestChannelNoSender(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(3 * time.Second)
		// channel <- "Rasyidev Pro"
		fmt.Println("Selesai mengirim data ke dalam channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
}

/*
$ go test -v -run TestChannelNoSender
=== RUN   TestChannelNoSender
Selesai mengirim data ke dalam channel
exit status 0xc000013a | Abort pake CTRL + C
FAIL    learn-go-goroutines     304.542s
*/

func TestChannelProper(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(3 * time.Second)
		channel <- "Rasyidev Pro"
		fmt.Println("Selesai mengirim data ke dalam channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
}

/*
$ go test -v -run TestChannelProper
=== RUN   TestChannelProper
Rasyidev Pro
Selesai mengirim data ke dalam channel
--- PASS: TestChannelProper (6.02s)
PASS
ok      learn-go-goroutines     6.053s
*/
