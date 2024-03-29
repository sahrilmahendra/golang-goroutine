package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	group := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
