package main

import "fmt"

func rangeSample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println(i, v)
	}
}