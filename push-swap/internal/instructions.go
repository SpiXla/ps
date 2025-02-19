package stack


func (a *Stack) Pa(b *Stack) {
	node := b.Pop()
	a.Push(node.Value)
}

func (b *Stack) Pb(a *Stack) {
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

func Ss(a, b *Stack) {
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

func Rr(a, b *Stack) {
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

func Rrr(a, b *Stack) {
	a.Rra()
	b.Rrb()
}
