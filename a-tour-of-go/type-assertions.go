package main

import "fmt"

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0, false

	f = i.(float64)
	fmt.Println(f)

}