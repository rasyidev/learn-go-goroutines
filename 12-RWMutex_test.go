package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Rekening struct {
	RWMutex sync.RWMutex
	Saldo   int
}

func (r *Rekening) TambahSaldo(s int) {
	r.RWMutex.Lock()
	r.Saldo += s
	r.RWMutex.Unlock()
}

func (r *Rekening) CekSaldo() int {
	r.RWMutex.RLock()
	saldo := r.Saldo
	r.RWMutex.RUnlock()
	return saldo
}

func TestRWMutex(t *testing.T) {
	rek := Rekening{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				rek.TambahSaldo(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Saldo saat ini:", rek.Saldo)
}

/*
$ go test -v -run TestRWMutex
=== RUN   TestRWMutex
Saldo saat ini: 100000
--- PASS: TestRWMutex (5.02s)
PASS
ok      learn-go-goroutines     5.077s
*/
