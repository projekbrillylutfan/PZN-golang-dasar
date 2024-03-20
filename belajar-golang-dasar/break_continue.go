package main

import "fmt"

func main() {
	// kode break (langsung dihentikan)
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
		
	// 	fmt.Println("perulangan ke :", i)
	// }

	// kode continue (di skip jika tidak memenuhi kondisi)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke: ", i)
	}
}