package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOnce)
		}()
	}
	group.Wait()
	fmt.Println(counter)
}
