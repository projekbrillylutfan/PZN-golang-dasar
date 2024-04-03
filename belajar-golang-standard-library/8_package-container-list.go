package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("ayam")
	data.PushBack("goreng")
	data.PushBack("telur")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}