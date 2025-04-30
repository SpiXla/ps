package stack

func Solver(a, b *Stack) []string {
	solution := []string{}
	mid := a.Length / 2
	for a.Length > 0 {
		if a.Length == 3 {
			solveThree(a, &solution)
			break
		}
		if a.Length == 5 {
			solveFive(a,b, &solution)
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
		if a.IsAsc(0, a.Length-1) && b.IsDsc(0, b.Length-1) {
			break
		}
	}
	for b.Length > 0 {
		solution = append(solution, "pa")
		a.Pa(b)
	}
	return solution
}

func solveThree(a *Stack, sol *[]string) {
	if a.Min() == 0 {
		if a.Max() == 2 {
			return
		}
		a.Sa()
		a.Ra()
		*sol = append(*sol, "sa", "ra")
	}
	if a.Min() == 1 {
		if a.Max() == 2 {
			a.Sa()
			*sol = append(*sol, "sa")
		} else {
			a.Ra()
			*sol = append(*sol, "ra")
		}
	}
	if a.Min() == 2 {
		if a.Max() == 1 {
			a.Rra()
			*sol = append(*sol, "rra")
		} else {
			a.Sa()
			a.Rra()
			*sol = append(*sol, "sa", "rra")
		}
	}
}

func solveFive(a,b *Stack, sol *[]string) {
	for a.Length > 3 {
		if a.Min() == 0 || a.Max() == 0 {
			b.Pb(a)
			*sol = append(*sol, "pb")
		}else if a.Min() == 1 || a.Max() == 1 {
			a.Sa()
			b.Pb(a)
			*sol = append(*sol, "sa")
			*sol = append(*sol, "pb")
		} else if a.Min() == a.Length || a.Max() == a.Length {
			a.Rra()
			b.Pb(a)
			*sol = append(*sol, "rra")
			*sol = append(*sol, "pb")
		}else {
			a.Rra()
			*sol = append(*sol, "rra")
		}
	}
	solveThree(a,sol)
	emptyB(a,b,sol)

}

func emptyB(a,b *Stack,sol *[]string) {
	a.Pa(b)
	*sol = append(*sol, "pa")
	if a.Max() == 0 {
		a.Ra()
		*sol = append(*sol, "ra")
	}
	a.Pa(b)
	*sol = append(*sol, "pa")
	if a.Max() == 0 {
		a.Ra()
		*sol = append(*sol, "ra")
	}
}