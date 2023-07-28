package main

import "fmt"

type User struct {
	name string
}

func (s User) cal(weight float64, height float64) float64 {
	result := weight / height / height * 10000
	return result
}

func main() {
	me := User{name: "yuu"}
	ibm := me.cal(59, 170)

	fmt.Println(ibm)
}