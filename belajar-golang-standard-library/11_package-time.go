package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Println(now.Local())

	// utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// fmt.Println(utc.Local())

	// parse, _ := time.Parse(time.RFC3339, "2020-01-01")
	// fmt.Println(parse)

	// time duration

	var duration1 time.Duration = time.Second * 10
	var duration2 time.Duration = time.Millisecond * 100
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration: %d\n", duration3 )
}
