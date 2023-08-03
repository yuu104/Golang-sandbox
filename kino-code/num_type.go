package main

import (
	"fmt"
	"reflect"
)

func numType() {
	num01 := 123
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(num01, reflect.TypeOf(num01))
	fmt.Println(num02, reflect.TypeOf(num02))
	fmt.Println(num03, reflect.TypeOf(num03))
	fmt.Println(num04, reflect.TypeOf(num04))
}