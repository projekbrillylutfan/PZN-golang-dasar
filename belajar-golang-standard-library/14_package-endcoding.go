package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	dataString := "ayam"
	var encode = base64.StdEncoding.EncodeToString([]byte(dataString))
	fmt.Println(encode)

	var decode, err = base64.StdEncoding.DecodeString(encode)
	if err == nil {
		fmt.Println(string(decode))
	} else {
		fmt.Println(err.Error())
	}
}