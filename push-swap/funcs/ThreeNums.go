package funcs

import "fmt"

// ThreeNums sorts an array of exactly three numbers
func ThreeNums(s []int) {
	if len(s) != 3 {
		fmt.Println("Error: ThreeNums expects exactly three numbers")
		return
	}

	first := s[0]
	second := s[1]
	last := s[2]



	// Sorting logic
	if LT(first, second) && LT(first, last) && GT(second, last) {
		SwapA() // 1 3 2 -> 1 2 3
		RotateA()
	}
	if GT(first, second) && GT(first, last) && GT(second, last) {
		SwapA() // 3 2 1 -> 1 2 3
		ReverseRotateA()
	}
	if GT(first, second) && LT(first, last) && LT(second, last) {
		ReverseRotateA() // 3 1 2 -> 1 2 3
	}
	if GT(first, second) && GT(first, last) && LT(second, last) {
		RotateA() // 2 1 3 -> 1 2 3
	}
	if LT(first, second) && GT(first,last) && GT(second,last) {
		ReverseRotateA() // 2 3 1 -> 1 2 3
	}

}

// Comparison functions
func GT(num1 int, num2 int) bool {
	return num1 > num2
}

func LT(num1 int, num2 int) bool {
	return num1 < num2
}
