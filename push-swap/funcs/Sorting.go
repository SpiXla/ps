package funcs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sorting() {
	arg1 := os.Args[1]
	parts := strings.Fields(arg1)

	for _, part := range parts {
		in, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error")
			return
		}
		if !IsUnique(in) {
			fmt.Println("Error")
			return
		}
		StackA = append(StackA, in)
	}

	if len(StackA) == 0 {
		fmt.Println("No numbers to sort.")
		return
	}

	for !IsSorted(StackA) {
		first := StackA[0]
		second := StackA[1]
		// last := StackA[len(StackA)-1]

		if len(StackA) == 3 {
			ThreeNums(StackA)
		}
		if len(StackA) < 3 && len(StackA) > 0 {
			if first > second {
				SwapA()
			}
		}
		// 1 2 3 4
		// 2 1 3 4
		//
		if len(StackA) > 3 {
			maxValue := StackA[0]

			for _, num := range StackA {
				if num > maxValue {
					maxValue = num
				}
			}
			
			for _, r := range StackA {
				if r < maxValue {
					PushB()
				}
			} 
		}
		
		if len(StackB) != 0 {
			PushA()
		}
	}
}

func IsUnique(in int) bool {
	for _, r := range StackA {
		if in == r {
			return false
		}
	}
	return true
}

func IsSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
