package learngogoroutines

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value float64, group *sync.WaitGroup) {
	data.Store(value, math.Pow(value, 2))
}

func TestAddToMap(t *testing.T) {
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(&data, float64(i), &group)
	}

	data.Range(func(key, value any) bool {
		fmt.Printf("%v:%v\n", key, value)
		return true
	})
}
