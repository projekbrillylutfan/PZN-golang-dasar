package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{
		"eko",
		"budi",
		"joko",
	}
	values := []int{
		100,
		200,
		300,
	}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "eko"))
	fmt.Println(slices.Index(names, "budi"))
}