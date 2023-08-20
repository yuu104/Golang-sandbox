package main

import (
	"fmt"
	"math"
)

type I2 interface {
	M()
}

type T2 struct {
	S string
}
func (t *T2) M() {
	fmt.Println(t.S)
}

type F float64
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var i I2

	i = &T2{"Hello"}
	describe(i) // (&{Hello}, main.T2)
	i.M() // Hello

	i = F(math.Pi)
	describe(i) // (3.14...., main.F)
	i.M() // 3.14.....
}