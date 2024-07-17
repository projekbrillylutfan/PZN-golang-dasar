package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RuneHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutine(t *testing.T) {
	go RuneHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}