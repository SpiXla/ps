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

func (s *Stack) Print() {
	pointer := s.Head
	for pointer != nil {
		fmt.Println(pointer.Value)
		pointer = pointer.Next
	}
}

// instruction

func (a *Stack) Pa(b Stack) {
	node := b.Pop()
	a.Push(node.Value)
}

func (b *Stack) Pb(a Stack) {
	node := a.Pop()
	b.Push(node.Value)
}

func (a *Stack) Sa() {
	third := a.Head.Next.Next
	head := a.Head.Next
	a.Head.Next.Next = a.Head
	a.Head.Next = third
	a.Head = head
}

func (b *Stack) Sb() {
	third := b.Head.Next.Next
	head := b.Head.Next
	b.Head.Next.Next = b.Head
	b.Head.Next = third
	b.Head = head
}

func Ss(a, b Stack) {
	b.Sb()
	a.Sa()
}

func (a *Stack) Ra() {
	first := a.Pop()
	node := a.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = first
}

func (b *Stack) Rb() {
	first := b.Pop()
	node := b.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = first
}

func Rr(a, b Stack) {
	b.Rb()
	a.Ra()
}
