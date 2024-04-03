package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}

	parseInt, err := strconv.ParseInt("1000", 10, 64)
	if err == nil {
		fmt.Println(parseInt)
	} else {
		fmt.Println("error", err.Error())
	}

	atoi, err := strconv.Atoi("200")
	if err == nil {
		fmt.Println(atoi)
	} else {
		fmt.Println("error", err.Error())
	}

	binary := strconv.FormatInt(1000, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(1000)

	fmt.Println(stringInt)
}