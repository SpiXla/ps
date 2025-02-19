package stack

import "fmt"

func Solver(a, b *Stack) {
	mid := a.length / 2
	// pointer := a.Head
	for a.length > mid {
		_, index := a.Min()
		if index == 0 {
			b.Pb(a)
			continue
		}
		if index == 1 {
			a.Sa()
			b.Pb(a)
			continue
		}
		if index <= mid {
			a.Ra()
			continue
		}
		if index > mid {
			a.Rra()
			continue
		}

	}

	

	fmt.Println("hthhghghhh")
	a.Print()
	fmt.Println("hthhghghhh")
	b.Print()
}
