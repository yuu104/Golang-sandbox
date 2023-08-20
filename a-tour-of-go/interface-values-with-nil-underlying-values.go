package main

import "fmt"

type I3 interface {
	M()
}

type T3 struct {
	S string
}
func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe2(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValuesWithNilUnderlyigValues() {
	var i I3

	var t *T3
	i = t
	describe2(i) // (<nil>, *main.T)
	i.M() // <nil>

	i = &T3{"hello"}
	describe2(i) // (&{hello}, *main.T)
	i.M() // hello
}