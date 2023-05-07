package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // semicolons can be omitted
		sum += sum
	}
	fmt.Println(sum)
}
