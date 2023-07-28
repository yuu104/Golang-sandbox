package main

import "fmt"

type User struct {
	gender string
	age int
}

func main() {
	var user1 User
	user1.gender = "male"
	user1.age = 20

	user2 := User{gender: "male", age: 20}

	fmt.Println(user1)
	fmt.Println(user2)
}