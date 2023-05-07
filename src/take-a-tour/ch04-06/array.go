package main

import "fmt"

func main() {
	var a [2]string // declaration without initialization
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // declare array primes and then initialize with a literal
	fmt.Println(primes)
}
