package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Head   *Node
	length int
}

func (s *Stack) Push(value int) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if s.Head == nil {
		s.Head = &newNode
		s.length++
	} else {
		newNode.Next = s.Head
		s.Head = &newNode
		s.length++
	}
}

func (s *Stack) Pop() *Node {
	if s.Head == nil {
		return nil
	}
	node := s.Head
	ptr := s.Head.Next
	s.Head.Next = nil
	s.Head = ptr
	s.length--
	return node
}

func (s *Stack) Min() (*Node, int) {
	pointer := s.Head
	min := pointer
	index := 0
	midx := 0
	for pointer != nil {
		if pointer.Value < min.Value {
			min = pointer
			midx = index
		}
		pointer = pointer.Next
		index++
	}
	return min, midx
}

func (s *Stack) Max() (*Node, int) {
	pointer := s.Head
	max := pointer
	index := 0
	madx := 0
	for pointer != nil {
		if pointer.Value > max.Value {
			max = pointer
			madx = index
		}
		pointer = pointer.Next
		index++
	}
	return max, madx
}

func (s *Stack) Print() {
	pointer := s.Head
	for pointer != nil {
		fmt.Println(pointer.Value)
		pointer = pointer.Next
	}
}
