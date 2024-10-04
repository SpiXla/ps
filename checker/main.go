package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var StackA []int
var StackB []int

func main() {
	if len(os.Args) < 2 {
		return // No input provided
	}

	arg1 := os.Args[1]
	parts := strings.Fields(arg1)

	for _, part := range parts {
		in, err := strconv.Atoi(part)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		if !isUnique(in) {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		StackA = append(StackA, in)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		executeCommand(command)
	}

	if isSorted(StackA) && len(StackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func isUnique(num int) bool {
	for _, n := range StackA {
		if n == num {
			return false
		}
	}
	return true
}

func isSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}

func executeCommand(command string) {
	switch command {
	case "sa":
		swapA()
	case "sb":
		swapB()
	case "pa":
		pushA()
	case "pb":
		pushB()
	case "ra":
		rotateA()
	case "rb":
		rotateB()
	case "rra":
		reverseRotateA()
	case "rrb":
		reverseRotateB()
	default:
		fmt.Fprintln(os.Stderr, "Error")
	}
}

// Implementing stack operations for checker
// Similar to `push-swap`, you can reuse some of the functions
func swapA() {
	if len(StackA) < 2 {
		return
	}
	StackA[0], StackA[1] = StackA[1], StackA[0]
}

func swapB() {
	if len(StackB) < 2 {
		return
	}
	StackB[0], StackB[1] = StackB[1], StackB[0]
}

func pushA() {
	if len(StackB) == 0 {
		return
	}
	StackA = append(StackA, StackB[len(StackB)-1])
	StackB = StackB[:len(StackB)-1]
}

func pushB() {
	if len(StackA) == 0 {
		return
	}
	StackB = append(StackB, StackA[len(StackA)-1])
	StackA = StackA[:len(StackA)-1]
}

func rotateA() {
	if len(StackA) == 0 {
		return
	}
	first := StackA[0]
	StackA = StackA[1:]
	StackA = append(StackA, first)
}

func rotateB() {
	if len(StackB) == 0 {
		return
	}
	first := StackB[0]
	StackB = StackB[1:]
	StackB = append(StackB, first)
}

func reverseRotateA() {
	if len(StackA) == 0 {
		return
	}
	last := StackA[len(StackA)-1]
	StackA = StackA[:len(StackA)-1]
	StackA = append([]int{last}, StackA...)
}

func reverseRotateB() {
	if len(StackB) == 0 {
		return
	}
	last := StackB[len(StackB)-1]
	StackB = StackB[:len(StackB)-1]
	StackB = append([]int{last}, StackB...)
}
