package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)

		if math.Sqrt(x) == z {
			break
		}
	}

	return z, nil
}

func exerciseEroors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
