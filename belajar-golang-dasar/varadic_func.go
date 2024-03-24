package main

import "fmt"

// di sini variabel numbers bisa dikatakan adalah variabel argumens(varargs)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	// jika data param nya berbentuk slice
	numbers := []int{10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}