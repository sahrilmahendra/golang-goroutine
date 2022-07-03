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

func ChannelOnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Sahril Mahendra"
	fmt.Println("Data telah dikirim")
}

func ChannelOnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Data diterima,", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go ChannelOnlyIn(channel)
	go ChannelOnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Moch."
		channel <- "Syahryil"
		channel <- "Mahendra"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}
