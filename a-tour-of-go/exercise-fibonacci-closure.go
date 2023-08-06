package main

import "fmt"

func fibonacci() func() int {
	prev1 := 0
	prev2 := 0

	return func() int {
		returnNum := prev1 + prev2
		prev2 = prev1
		prev1 = returnNum
		if prev1 == 0 && prev2 == 0 {
			prev2 = 1
		}
		
		return returnNum
	}
}

func execiseFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

