package main

import "fmt"


func endApp() {
	fmt.Println("aplikasi selesai")
	message := recover()
	fmt.Println("terjadi error di ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}

	// contoh recover salah
	// message := recover()
	// fmt.Println("terjadi error di ", message)
}

func main() {
	runApp(true)
}