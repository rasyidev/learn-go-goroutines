package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TheRealAsynchronous(group *sync.WaitGroup) {
	/*
		untuk jaga jaga jika ada error pada program
		Kalau group.Done() tidak ada, maka WaitGroup akan
		menunggu terus, sampai deadlock karena tidak selesai selesai
	*/
	defer group.Done()

	group.Add(1)
	fmt.Println("Proses selesai")
	time.Sleep(2 * time.Second)
}

func TestTheRealAsynchronous(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go TheRealAsynchronous(group)
	}
	group.Wait()
	fmt.Println("Selesai")
}

/*
Proses selesai
.....
Proses selesai
Selesai
--- PASS: TestTheRealAsynchronous (2.06s)
PASS
ok      learn-go-goroutines     2.102s
*/
