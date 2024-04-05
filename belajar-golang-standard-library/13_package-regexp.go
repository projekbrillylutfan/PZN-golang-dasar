package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("kang"))
	fmt.Println(regex.MatchString("edelweis"))

	fmt.Println(regex.FindAllString("eko edelweis ayam dus aja kita ejo eMo", 10))
}