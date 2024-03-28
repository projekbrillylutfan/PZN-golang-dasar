package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "set host")
	username := flag.String("username", "admin", "set username")
	password := flag.String("password", "admin", "set password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}