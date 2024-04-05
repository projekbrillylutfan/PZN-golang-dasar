package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// csv read
	csvString := "eko, kurniawan, khannedy\n" +
		"budi, setiawan, bambang\n" +
		"joko, widodo,"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
