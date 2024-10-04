package main

import (
	"fmt"
	"os"

	"funcs/push-swap/funcs"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	funcs.Sorting()
	for _, r := range funcs.Instructions {
		fmt.Println(r)
	}

	// funcs.PushB()

	fmt.Println("stackA:",funcs.StackA)
	fmt.Println("stackB:",funcs.StackB)
}
