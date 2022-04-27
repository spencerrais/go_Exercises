package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 11; i++ {
		if z == (z*z-x)/(2*z) || (z*z-x)/(2*z) < .0000000000001 && i > 1 {
			return z
		} else {
			z -= (z*z - x) / (2 * z)
			fmt.Println(z)
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
