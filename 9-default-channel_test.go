package learngogoroutines

import (
	"fmt"
	"testing"
)

func TestDefaultSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}

/*
$ go test -v -run TestDefaultSelectChannel
Menunggu data
Menunggu data
.....
Menunggu data
Menunggu data
Data dari channel 1 Rasyidev
Data dari channel 2 Rasyidev
--- PASS: TestDefaultSelectChannel (2.01s)
PASS
ok      learn-go-goroutines     2.056s
*/
