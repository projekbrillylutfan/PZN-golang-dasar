package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i:= 0; i < 20; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	} ()

	for data:= range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
	
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}


}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}


}
