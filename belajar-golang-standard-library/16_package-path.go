package main

import (
	"fmt"
	// "path"
	"path/filepath"
)

func main() {
	// contoh package path
	// fmt.Println(path.Dir("hello/world.go"))
	// fmt.Println(path.Base("hello/world.go"))
	// fmt.Println(path.Ext("hello/world.go"))
	// fmt.Println(path.Join("hello", "world.go"))

	// contoh package path/filepath
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world.go"))
}