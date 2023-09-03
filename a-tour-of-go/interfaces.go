package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}
func (t T) M() {
	fmt.Println(t.S)
}

type Test struct {
	S string
}
func (test *Test) M() {
	fmt.Println(test.S)
}

func interfaces() {
	var i I = T{"hello"}
	i.M()
}