package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	// seharusnya 100000
	time.Sleep(5 * time.Second)
	fmt.Println("X saat ini:", x)
}

/* Run Pertama
$ go test -v -run TestRaceCondition
=== RUN   TestRaceCondition
X saat ini: 86872
--- PASS: TestRaceCondition (5.00s)
PASS
ok      learn-go-goroutines     5.046s
*/

/* Run Kedua
$ go test -v -run TestRaceCondition
=== RUN   TestRaceCondition
X saat ini: 95899
--- PASS: TestRaceCondition (5.00s)
PASS
ok      learn-go-goroutines     5.047s
*/
