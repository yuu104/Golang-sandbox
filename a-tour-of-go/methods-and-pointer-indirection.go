package main

import "fmt"

type VertxMethodPointer struct {
	X, Y float64
}

func (v *VertxMethodPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *VertxMethodPointer, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodsAndPointerIndirection() {
	v := VertxMethodPointer{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &VertxMethodPointer{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}