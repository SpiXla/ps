package main

import (
	stack "pushswap/internal"
)

func main() {
	s := stack.Stack{}
	s.Push(8)
	s.Push(5)
	s.Push(6)
	s.Push(3)
	s.Push(1)
	s.Push(2)

	s.Print()
}
