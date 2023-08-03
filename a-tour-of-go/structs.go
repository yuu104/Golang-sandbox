package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println(Vertex{1, 2})
}

func structFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func pointersToStructs() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, p)
}