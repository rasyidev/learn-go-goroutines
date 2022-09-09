package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type RekeningBSI struct {
	sync.Mutex // alternatif: Mutex sync.Mutex
	Nama       string
	Saldo      int
}

func (r *RekeningBSI) Lock() {
	r.Mutex.Lock()
}

func (r *RekeningBSI) Unlock() {
	r.Mutex.Unlock()
}

func (r *RekeningBSI) TambahSaldo(rp int) {
	// r.Lock()
	r.Saldo += rp
	// r.Unlock()
}

/*
Mentransfer dari rek a ke rek b dengan jumlah tertentu
from: rekening pengirim
to: rekening penerima
rp: Uang yang dikirim dalam Rupiah (rp)
*/
func Transfer(from *RekeningBSI, to *RekeningBSI, rp int) {
	from.Lock()
	fmt.Println("Lock", from.Nama)
	from.TambahSaldo(-rp)
	time.Sleep(2 * time.Second)

	to.Lock()
	fmt.Println("Lock", from.Nama)
	to.TambahSaldo(rp)
	time.Sleep(2 * time.Second)

	from.Unlock()
	to.Unlock()
}

func TestDeadlock(t *testing.T) {
	andi := RekeningBSI{
		Nama:  "Andi",
		Saldo: 1000000,
	}

	budi := RekeningBSI{
		Nama:  "Budi",
		Saldo: 1000000,
	}

	go Transfer(&andi, &budi, 100000)
	go Transfer(&budi, &andi, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("Saldo Andi:", andi.Saldo)
	fmt.Println("Saldo Budi:", budi.Saldo)
}

/*
=== RUN   TestDeadlock
Lock Budi | Lock, kurangi 200rb, Gak di-unlock - unlock padahal mau kirim 200rb ke Andi
Lock Andi | Lock, kurangi 100rb, Gak di-unlock - unlock padahal mau kirim 100rb ke Budi
Saldo Andi: 900000
Saldo Budi: 800000
--- PASS: TestDeadlock (10.01s)
PASS
ok      learn-go-goroutines     10.048s
*/
