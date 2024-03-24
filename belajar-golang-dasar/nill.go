package main

import "fmt"

// nill bentuk map
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	contoh := NewMap("")
	if contoh == nil {
		fmt.Println("map kosong")
	} else {
		fmt.Println(contoh)
	}
}