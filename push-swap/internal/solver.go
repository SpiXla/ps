package stack

func Solver(a, b *Stack) []string {
	solution := []string{}
	mid := a.Length / 2
	// pointer := a.Head
	for a.Length > 0 {
		if a.IsAsc(0, a.Length-1) && b.IsDsc(0, b.Length-1) {
			break
		}
		if a.Min() == 0 {
			b.Pb(a)
			solution = append(solution, "pb")
			continue
		}
		if a.Min() == 1 {
			if b.Max() == 1 {
				solution = append(solution, "ss")
				Ss(a, b)
			} else {
				solution = append(solution, "sa")
				a.Sa()
			}
			solution = append(solution, "pb")
			b.Pb(a)
			continue
		}
		if a.Min() <= mid {
			if a.Max() == 0 && b.Min() == 0 {
				solution = append(solution, "rr")
				Rr(a, b)
			} else {
			solution = append(solution, "ra")
				a.Ra()
			}
			continue
		}
		if a.Min() > mid {		
			if a.Min() == 0 && b.Max() == 0 {
				solution = append(solution, "rrr")
				Rrr(a, b)
			} else {
				solution = append(solution, "rra")
				a.Rra()
			}
			continue
		}

	}
	for b.Length > 0 {
		solution = append(solution, "pa")
		a.Pa(b)
	}

	return solution
}
