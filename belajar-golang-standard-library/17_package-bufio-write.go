package main

import (
	"bufio"
	"os"
)

func main() {
	// bufio reader
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello World\nayam goren\n cihuy")
	writer.WriteString("\nselamat berbuka puasa")
	writer.Flush()
}