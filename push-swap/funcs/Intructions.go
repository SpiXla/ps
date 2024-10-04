package funcs

func PushA() {
	StackA = append(StackA, StackB[0])
	StackB = StackB[1:]
	Instructions = append(Instructions, "pa")
}

func PushB() {
	StackB = append(StackB, StackA[0])
	StackA = StackA[1:]
	Instructions = append(Instructions, "pb")
}

func SwapA() {
	StackA[0], StackA[1] = StackA[1], StackA[0]
	Instructions = append(Instructions, "sa")
}

func SwapB() {
	StackB[0], StackB[1] = StackB[1], StackB[0]
	Instructions = append(Instructions, "sb")
}

func SwapBoth() {
	SwapA()
	SwapB()
	Instructions = append(Instructions, "ss")
}

func RotateA() {
	StackA = append(StackA[1:], StackA[0])
	Instructions = append(Instructions, "ra")
}

func RotateB() {
	StackB = append(StackB[1:], StackB[0])
	Instructions = append(Instructions, "rb")
}

func RotateBoth() {
	RotateA()
	RotateB()
	Instructions = append(Instructions, "rr")
}

func ReverseRotateA() {
	StackA = append(StackA[len(StackA)-1:], StackA[0:len(StackA)-1]...)
	Instructions = append(Instructions, "rra")
}

func ReverseRotateB() {
	StackB = append(StackB[len(StackB)-1:], StackB[0:len(StackB)-1]...)
	Instructions = append(Instructions, "rrb")
}

func ReverseRotateBoth() {
	ReverseRotateA()
	ReverseRotateB()
	Instructions = append(Instructions, "rrr")
}
