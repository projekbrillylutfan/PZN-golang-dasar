package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Message 1")
	pool.Put("Message 2")
	pool.Put("Message 3")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}