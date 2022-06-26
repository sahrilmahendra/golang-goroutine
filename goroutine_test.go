package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunGoroutine() {
	fmt.Println("Hello World")
}

// membuat goroutine
func TestGoroutine(t *testing.T) {
	go RunGoroutine()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

// membuktikan goroutine sangat ringan
func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i <= 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
