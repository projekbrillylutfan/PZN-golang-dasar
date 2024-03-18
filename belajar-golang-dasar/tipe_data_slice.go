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
	// days := [...]string {
	// 	"senin",
	// 	"selasa",
	// 	"rabu",
	// 	"kamis",
	// 	"jumat",
	// 	"sabtu",
	// 	"minggu",
	// }

	// daySlice := days[5:]
	// daySlice[0] = "Sabtu Baru"
	// daySlice[1] = "Minggu Baru"
	// fmt.Println(days)

	// daySlice2 := append(daySlice, "libur")
	// daySlice2[0] = "yeah"
	// daySlice2[0] = "sabtu lama"
	// fmt.Println(daySlice2)
	// fmt.Println(days)
	// fmt.Println(daySlice)

	// bikin slice baru
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "brilly"
	// newSlice[1] = "lutfan"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// newSlice2 := append(newSlice, "qasthori")
	// fmt.Println(newSlice2)
	// fmt.Println(len(newSlice2))
	// fmt.Println(cap(newSlice2))

	// copy slice
	// fromSlice := days[:]
	// toSlice := make([]string, len(fromSlice), cap(fromSlice))

	// copy(toSlice, fromSlice)

	// fmt.Println(toSlice)

	//perbedaan array dan slice

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}