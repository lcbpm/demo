package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialized")
}

func Test(t *testing.T) {
	for i := 0; i < 10; i++ {
		go once.Do(initialize)
	}
	time.Sleep(time.Second)
}
