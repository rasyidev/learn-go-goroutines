package learngogoroutines

import (
	"fmt"
	"testing"
)

func TestSelectChannelNoLoop(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	select {
	case data := <-ch1:
		fmt.Println("Data dari channel 1", data)
	case data := <-ch2:
		fmt.Println("Data dari channel 2", data)
	}
}

/*
$ go test -v -run TestSelectChannelNoLoop
=== RUN   TestSelectChannelNoLoop
Data dari channel 2 Rasyidev
panic: send on closed channel
*/

func TestSelectChannelWithLoop(t *testing.T) {
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
		}

		if counter == 2 {
			break
		}
	}
}

/*
$ go test -v -run TestSelectChannelWithLoop
=== RUN   TestSelectChannelWithLoop
Data dari channel 2 Rasyidev
Data dari channel 1 Rasyidev
--- PASS: TestSelectChannelWithLoop (2.01s)
PASS
ok      learn-go-goroutines     2.049s
*/
