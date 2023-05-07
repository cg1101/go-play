package main

import (
	"fmt"
	//"strings"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		//fmt.Println("walk left")
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		//fmt.Println("walk right")
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	seq1 := make([]int, 10)
	seq2 := make([]int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		seq1 = append(seq1, <-ch1)
	}
	for j := 0; j < 10; j++ {
		seq2 = append(seq2, <-ch2)
	}
	rs1 := fmt.Sprintf("%q", seq1)
	rs2 := fmt.Sprintf("%q", seq2)

	return rs1 == rs2
}

func main() {
	t1a := tree.New(1)
	t1b := tree.New(1)
	t2 := tree.New(2)
	fmt.Printf("Same(t1a, t1b) -> %v\n", Same(t1a, t1b))
	fmt.Printf("Same(t1a, t2) -> %v\n", Same(t1a, t2))
}
