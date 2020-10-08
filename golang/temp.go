package main

import (
	"fmt"
	"math"
)

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func main() {
	n := 4
	fmt.Println(maxDepth(n))
	m := 4.0
	fmt.Println(2 * math.Ceil(math.Log10(m+1.0))) // 2*ceil(lg(n+1)).
}
