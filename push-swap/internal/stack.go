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

func (s *Stack) MaxMin() (*Node, int, *Node,int) {
	pointer := s.Head
	min := pointer
	max := pointer
	index := 0
	midx := 0
	madx := 0
	for pointer != nil {
		if pointer.Value < min.Value {
			min = pointer
			midx = index
		}
		if pointer.Value > max.Value {
			max = pointer
			madx = index
		}
		pointer = pointer.Next
		index++
	}
	return min,midx,max, madx
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

func (a *Stack) Rra() {
	node := a.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	last := node.Next
	last.Next = a.Head
	node.Next = nil
	a.Head = last
}

func (b *Stack) Rrb() {
	node := b.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	last := node.Next
	last.Next = b.Head
	node.Next = nil
	b.Head = last
}

func Rrr(a, b Stack) {
	a.Rra()
	b.Rrb()
}
