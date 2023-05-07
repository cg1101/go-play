package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(x int) int {
	known := []int{0, 1}
	return func(x int) int {
		for i := len(known); i <= x; i++ {
			v := known[i-1] + known[i-2]
			known = append(known, v)
		}
		return known[x]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
