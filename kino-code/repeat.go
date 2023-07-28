package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	l := 1
	for l <= 10 {
		fmt.Println(l)
		l++
	}

	fmt.Println("---------------")

	n := 1
	for {
		if n > 10 {
			break
		}
		fmt.Println(n)
		n++
	}
}