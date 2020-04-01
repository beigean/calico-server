// ref. https://go-tour-jp.appspot.com/methods/20
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z, z_prev := 1.0, x
	const accuracy = 1.0E-10
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if (z_prev - z) < accuracy {
			break
		}
		z_prev = z
	}
	return z, nil
}

func main() {
	target := float64(-3648)
	if result, err := Sqrt(target); err == nil {
		fmt.Println("self", result)
	} else {
		fmt.Println(err)
	}

	fmt.Println("math", math.Sqrt(target))
}
