package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunGoroutine() {
	fmt.Println("Hello World")
}

func TestGoroutine(t *testing.T) {
	go RunGoroutine()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}
