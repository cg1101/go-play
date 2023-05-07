package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	v := float64(e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", v)
}

func Sqrt(x float64) (float64, error) {
	v := x
	if x < 0 {
		return 0, ErrNegativeSqrt(v)
		//return 2, nil
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
