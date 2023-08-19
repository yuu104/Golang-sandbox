package main

import (
	"fmt"
	"math"
)

type VertxMethodPointer2 struct {
	X, Y float64
}

func (v VertxMethodPointer2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v VertxMethodPointer2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodsAndPointerIndirection2() {
	v := VertxMethodPointer2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertxMethodPointer2{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}