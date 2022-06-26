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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Moch. Syahryil Mahendra"
	fmt.Println("Data telah dikirim")
}

func TestChannelAsParameter(t *testing.T) {
	testChannel := make(chan string)
	defer close(testChannel)

	go GiveMeResponse(testChannel)

	data := <-testChannel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
}
