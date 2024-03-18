package main

import "fmt"

func main() {
	// var names [3]string
	// names[0] = "brilly"
	// names[1] = "lutfan"
	// names[2] = "qasthari"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	//membuat array langsung
	// var values = [5]int{
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// 	5,
	// }

	// fmt.Println(values)

	//function array
	var values = [...]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 10
	fmt.Println(values)
}