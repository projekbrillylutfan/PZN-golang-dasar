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
