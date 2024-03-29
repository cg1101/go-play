package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var wc = make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w] = int(wc[w]) + 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
