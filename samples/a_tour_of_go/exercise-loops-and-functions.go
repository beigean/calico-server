// ref. https://go-tour-jp.appspot.com/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
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
