package main

import "fmt"

func main() {
	ch := make(chan int, 2) // create a channel for int, buffer size is 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
