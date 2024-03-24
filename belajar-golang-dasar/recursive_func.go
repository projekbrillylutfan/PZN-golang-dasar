package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func factorLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorLoop(10))
	fmt.Println(factorialRecursive(10))
}