package main

import "fmt"

type User struct {
	gender string
	age int
}

func (s User) calIbm(weight float64, height float64) float64 {
	result := weight / height / height * 10000
	return result
}

func structure() {
	var user1 User
	user1.gender = "male"
	user1.age = 20

	user2 := User{gender: "male", age: 20}

	user2Ibm := user2.calIbm(59, 170)

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user2Ibm)
}