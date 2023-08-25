package main

import "fmt"

type I4 interface {
	M()
}

func nilInterfaceValues() {
	var i I
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()
}
