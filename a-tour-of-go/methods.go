package main

import (
	"fmt"
	"math"
)

type VectorMethod struct {
	X, Y float64
}

func (v VectorMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method() {
	v := VectorMethod{3, 4}
	fmt.Println(v.Abs())
}