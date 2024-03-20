package main 

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke :", counter)
	// 	counter++
	// }

	// fmt.Println("sudah selesai")

	// for statment
	// for counter :=1; counter <=10; counter++ {
	// 	fmt.Println("perulangan ke :", counter)
	// }

	// fmt.Println("sudah selesai")

	// for range
	names := []string{
		"eko",
		"ayam",
		"jantan",
	}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for i, name := range names {
		fmt.Println("index", i, "=", name)
	}

	// misal tidak butuh index nya
	for _, name := range names {
		fmt.Println("hasil :", name)
	}
}