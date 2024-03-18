package main

import "fmt"

func main() {
	// names := [...]string{
	// 	"brilly",
	// 	"lutfan",
	// 	"qasthori",
	// 	"billy1",
	// 	"lutfan1",
	// 	"qasthori1",
	// }

	// slice1 := names[4:6]
	// fmt.Println(slice1)

	// slice2 := names[:3]
	// fmt.Println(slice2)

	// slice3 := names[3:]
	// fmt.Println(slice3)

	// slice4 := names[:]
	// fmt.Println(slice4)

	// //brakedown tipe data slice
	// var slice5 []string = names[:]
	// fmt.Println(slice5)

	// //len
	// fmt.Println("ini len : ",len(slice5))

	// append
	days := [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	daySlice := days[5:]
	daySlice[0] = "Sabtu Baru"
	daySlice[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice, "libur")
	daySlice2[0] = "yeah"
	fmt.Println(daySlice2)
	fmt.Println(days)

}