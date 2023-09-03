package main

import (
	"fmt"
	"io"
	"strings"
)

func reader2() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v len(b) = %v\n", n, err, len(b))
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// n = 8 err = <nil> len(b) = 8
// b[:n] = Hello, R
// n = 6 err = <nil> len(b) = 8
// b[:n] = eader!
// n = 0 err = EOF len(b) = 8
// b[:n] = 