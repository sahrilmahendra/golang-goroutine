package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var locker = &sync.Mutex{}
var cond = sync.NewCond(locker)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// // cond.Signal() --> sebagai tanda (flag) untuk menjalankan 1 goroutine
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()
	cond.Broadcast()

	group.Wait()
}
