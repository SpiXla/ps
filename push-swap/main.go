package main

import (
	"fmt"

	stack "pushswap/internal"
)

func main() {
	// n := []int{1, 10}

	a := stack.Stack{Name: "A"}
	b := stack.Stack{Name: "B"}
	// for a.Length < 100 {
	// 	r := int(rand.Float64() * 10)
	// 	a.Push(r * n[r%2])
	// }
	a.Push(3)
	a.Push(5)
	a.Push(6)
	a.Push(2)
	a.Push(1)

	
	// a.Push(1)
	// a.Push(12)
	// a.Push(17)
	// a.Push(14)
	// a.Push(11)

	a.Print()

	// a.Push(12)
	// a.Push(17)
	// a.Push(14)
	// a.Push(11)
	// a.Push(8)
	// a.Push(6)
	// a.Print()

	// a.Print()

	sol := stack.Solver(&a, &b)
	a.Print()
	fmt.Println(sol)
	fmt.Println("instructions used :", len(sol))
}
