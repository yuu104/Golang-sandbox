package main

import "fmt"

func relationalOperator() {
	x := 10
	y := 2
	z := 10

	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x >= z)
	fmt.Println(x == y)
	fmt.Println(x != y)

	a := 8
	b := 3

	fmt.Println(a >= 5 && b <= 10)
	fmt.Println(b <= 5 && b <= 10)
	fmt.Println(a == 3 || b == 3)
	fmt.Println(a == 1 || b == 1)
}