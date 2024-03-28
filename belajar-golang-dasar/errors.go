package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	pembagi, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println(pembagi)
	} else {
		fmt.Println(err.Error())
	}
	
}