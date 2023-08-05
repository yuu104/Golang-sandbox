package main

import (
	"fmt"
	"strings"
)

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s[]int = primes[1:4]
	fmt.Println(s)
}

// スライスの要素を変更すると、その元となる配列の要素も変更が反映される
func sliceAreLikeReferencesToArrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}

func sliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s) // []

	s = s[:4]
	printSlice(s) // [2, 3, 5, 7]

	s = s[2:]
	printSlice(s) // [5, 7]
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func creatingASliceWithMake() {
	a := make([]int, 5)
	fmt.Printf("%s, len=%d, cap=%d, %v\n", "a", len(a), cap(a), a)
	// a, 5, 5, [0, 0, 0, 0, 0]

	b := make([]int, 0, 5)
	fmt.Printf("%s, len=%d, cap=%d, %v\n", "b", len(b), cap(b), b)
	// b, 0, 5, []

	c := b[:2]
	fmt.Printf("%s, len=%d, cap=%d, %v\n", "c", len(c), cap(c), c)
	// c, 2, 5, [0, 0]

	d := c[2:5]
	fmt.Printf("%s, len=%d, cap=%d, %v\n", "d", len(d), cap(d), d)
	// d, 3, 3, [0, 0, 0]
}

func sliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingToASlice() {
	s := []int{0, 1}
	a := append(s, 2)
	a[0] = 9
	printSlice(a)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}