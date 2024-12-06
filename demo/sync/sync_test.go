package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once

var m map[int]int
var l []int64

func initialize() {
	fmt.Println("Initialized")
}

func Test(t *testing.T) {
	if m == nil || l == nil {

	}
	for i := 0; i < 10; i++ {
		go once.Do(initialize)
	}
	time.Sleep(time.Second)
}
