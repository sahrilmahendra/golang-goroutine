package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Sahril Mahendra"
		fmt.Println("Data telah dikirim ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
