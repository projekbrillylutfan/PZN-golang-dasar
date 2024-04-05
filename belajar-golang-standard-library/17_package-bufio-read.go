package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	// bufio reader
	input := strings.NewReader("Hello World\nayam goren\n cihuy")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}