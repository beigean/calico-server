// ref. https://go-tour-jp.appspot.com/methods/20
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: -2")
}

func Sqrt(x float64) (float64, error) {
	error err = nil
	if x < 0 {
		err = ErrNegativeSqrt(-2)
	}

	z, z_prev := 1.0, x
	const accuracy = 1.0E-10
	fmt.Println(accuracy)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if (z_prev - z) < accuracy {
			break
		}
		z_prev = z
	}
	return z
}

func main() {
	target := float64(3648)
	fmt.Println("self", Sqrt(target))
	fmt.Println("math", math.Sqrt(target))
}
