package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Hello Channel"
		fmt.Println("Selesai Mengirim Data ke channel")
	}()

	data:= <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	defer close(channel)
}

func GiveMeResponse (channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "ayam goreng"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}

func OnlyIn (channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "ayam goreng"
}

func OnlyOut (channel <-chan string) {
	data:= <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	go OnlyIn(channel)

	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "Hello"
		channel <- "World"
		channel <- "Go"
	} ()

	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	} ()
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}
