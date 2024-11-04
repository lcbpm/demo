package gc

import (
	"log"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 你的程序逻辑
	select {}
}
