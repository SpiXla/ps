package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Name   string
	Head   *Node
	Length int
}

func (s *Stack) Push(value int) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if s.Head == nil {
		s.Head = &newNode
		s.Length++
	} else {
		newNode.Next = s.Head
		s.Head = &newNode
		s.Length++
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
	s.Length--
	return node
}

func (s *Stack) Min() int {
	if s.Length == 0 {
		return -1
	}
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
	return midx
}

func (s *Stack) Max() int {
	if s.Length == 0 {
		return -1
	}
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
	return madx
}

func (s *Stack) IsAsc(start, end int) bool {
	if s.Length <= 0 {
		return true
	}
	index := 0
	reach := false
	pointer := s.Head
	for pointer != nil {
		if index == end {
			break
		}
		if index == start {
			reach = true
		}
		if reach {
			if pointer.Value > pointer.Next.Value {
				return false
			}
		}
		pointer = pointer.Next
		index++
	}
	return true
}

func (s *Stack) IsDsc(start, end int) bool {
	if s.Length <= 1 {
		return false
	}
	index := 0
	reach := false
	pointer := s.Head
	for pointer != nil {
		if index == end {
			break
		}
		if index == start {
			reach = true
		}
		if reach {
			if pointer.Value < pointer.Next.Value {
				return false
			}
		}
		pointer = pointer.Next
		index++
	}
	return true
}

func (s *Stack) Print() {
	fmt.Println("stack", s.Name)
	fmt.Println("=====")
	pointer := s.Head
	for pointer != nil {
		fmt.Println(pointer.Value)
		pointer = pointer.Next
	}
	fmt.Println("=====")
}

func (s *Stack) Get(from, to int) []int {
	elements := []int{}
	collect := false
	pointer := s.Head
	i := 0
	for pointer != nil {
		if i >= from {
			collect = true
		}
		if collect {
			elements = append(elements, pointer.Value)
		}
		if i == to {
			collect = false
		}
		pointer = pointer.Next
		i++
	}
	return elements
}
