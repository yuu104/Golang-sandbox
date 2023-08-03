package main

import (
	"fmt"
	"reflect"
)

func stringType() {
	var string_a string = "Hello, World!"
	string_b := "Hello, World!"

	fmt.Println(string_a, reflect.TypeOf(string_a))
	fmt.Println(string_b, reflect.TypeOf(string_b))
}