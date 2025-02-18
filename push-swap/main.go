package main

import (
	"fmt"
	stack "pushswap/internal"
)

func main() {
	n := stack.Stack{}
	n.Push(4)
	n.Push(2)
	n.Push(1)
	n.Print()
	fmt.Println("---------")
	n.Sa()
	n.Print()
	fmt.Println("---------")
	n.Ra()
	n.Print()
}
