package main

import "fmt"

//contoh defer
func logging() {
	fmt.Println("selesai")
}

func runAplication() {
	defer logging()
	fmt.Println("mulai")
}

func main() {
	runAplication()
}
