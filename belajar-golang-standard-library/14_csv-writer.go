package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// csv write
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"budi", "setiawan", "bambang"})
	_ = writer.Write([]string{"joko", "widodo", ""})

	writer.Flush()
}