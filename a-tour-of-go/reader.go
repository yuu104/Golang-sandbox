package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data string
	pos  int
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

func NewStringReader(s string) *StringReader {
	return &StringReader{data: s}
}

func reader2() {
	// reader := NewStringReader("Hello, World!")
	// buffer := make([]byte, 6)

	// for {
	// 	n, err := reader.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	// 	fmt.Println(string(buffer))
	// }
	arr := []int{1, 2, 3}
	arr2 := make([]int, 6)
	copy(arr2, arr)
	fmt.Println(arr2)
	copy(arr2, []int{3, 2, 1})
	fmt.Println(arr2)
}
