package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var stillBig = true
	for counter := 0; stillBig && counter < 200; counter++ {
		delta := (z*z - x) / (2 * z)
		stillBig = math.Abs(delta) > 1.e-10
		fmt.Println("counter=", counter, "z=", z, "delta=", delta, "big?", stillBig)
		z -= delta
	}
	return z

}

func main() {
	fmt.Println("Square root of 2 is", Sqrt(2))
	fmt.Println("result from math.Sqrt", math.Sqrt(2))
	fmt.Println("math.abs(-2)", math.Abs(-2))
}
