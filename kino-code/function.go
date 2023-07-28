package main

import "fmt"

func main() {
	fmt.Println(cal(10, 5))
}

func cal(x int, y int) int {
	sum := x + y
	return sum
}