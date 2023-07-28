package main

import "fmt"

func main() {
	a := [3]string{"sato", "suzuki", "takahashi"}

	fmt.Println(a)
	fmt.Println(a[0], a[1], a[2])

	a[1] = "tanaka"
	fmt.Print(a)

	b := [...]string{"sato", "suzuki", "tanaka"}
	fmt.Println(b)

	c := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}
	fmt.Println(c)
}