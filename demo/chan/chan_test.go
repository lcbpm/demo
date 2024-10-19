package _chan

import (
	"fmt"
	"testing"
	"time"
)

type chanService struct {
}

func NewChanService() *chanService { return new(chanService) }

func (c *chanService) chan1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "chan 1"
}

func (c *chanService) chan2(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "chan 2"
}

func (c *chanService) handle() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go c.chan1(ch1)
	go c.chan2(ch2)

	//阻塞
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)

	}

}

func (c *chanService) handle2() {
	
}

func TestChanService(t *testing.T) {
	NewChanService().handle()
}
