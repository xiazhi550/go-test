package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z float64 = 1
	var d float64
	for i := 0; i < 10; i++ {
		d = (z*z - x) / (2 * z)
		z -= d
		fmt.Println(z)
		if math.Abs(d) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
