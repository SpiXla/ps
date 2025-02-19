package main

import (
	"fmt"
	"math/rand"
	stack "pushswap/internal"
)

func main() {
	n := []int{1, 10}

	a := stack.Stack{}
	for a.Length < 100 {
		r := int(rand.Float64() * 10)
		a.Push(r * n[r%2])
	}
	// a.Push(8)
	// a.Push(5)
	// a.Push(6)
	// a.Push(3)
	// a.Push(1)
	// a.Push(2)
	
	b := stack.Stack{}
	sol := stack.Solver(&a, &b)
	fmt.Println(sol)
	a.Print()
}
