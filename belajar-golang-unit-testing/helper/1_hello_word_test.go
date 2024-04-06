package helper

import (
	"fmt"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWord(t *testing.T) {
// 	result := HelloWord("brilly")
// 	if result != "Hello brilly" {
// 		t.Fail()
// 	}

// 	fmt.Println("Test success func TestHelloWord")
// }

// func TestHelloWordBjir(t *testing.T) {
// 	result := HelloWord("bjir")
// 	if result != "Hello bjir" {
// 		t.FailNow()
// 	}

// 	fmt.Println("Test success func TestHelloWordBjir")
// }

// func TestHelloWord(t *testing.T) {
// 	result := HelloWord("brilly")
// 	if result != "Hello brilly" {
// 		t.Error("harus hello brilly")
// 	}

// 	fmt.Println("Test success func TestHelloWord")
// }

// func TestHelloWordBjir(t *testing.T) {
// 	result := HelloWord("bjir")
// 	if result != "Hello bjir" {
// 		t.Fatal("harus hello bjir")
// 	}

// 	fmt.Println("Test success func TestHelloWordBjir")
// }

// menggunakan assertion testify

// func TestHelloWordAssertion(t *testing.T) {
// 	result := HelloWord("brilly")
// 	assert.Equal(t, "Hello brilly", result, "hasil harus Hello brilly")
// 	fmt.Println("dieksekusi")
// }

// menggunakan require testify
func TestHelloWordRequire(t *testing.T) {
	result := HelloWord("brilly")
	require.Equal(t, "Hello brilly", result, "hasil harus Hello brilly")
	fmt.Println("tidak dieksekusi")
}