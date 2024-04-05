package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"100"`
}

type Person struct {
	Name    string `required:"true" max:"100"`
	Address string `required:"true" max:"100"`
	Email   string `required:"true" max:"100"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(data interface{}) (result bool) {
	result = true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface()
			result = data != ""
			if !result {
				return false
			}
		}
	}
	
	return true
}

func main() {
	// readField(Sample{Name: "eko"})
	// readField(Person{Name: "eko"})

	person := Person {
		Name: "eko",
		Address: "",
		Email: "",
	}

	fmt.Println(IsValid(person))
}
