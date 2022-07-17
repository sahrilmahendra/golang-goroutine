package golanggoroutine

import (
	"fmt"
	"strconv"
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

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data ", data)
	}
	fmt.Println("DONE")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// default select channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
