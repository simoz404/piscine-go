package piscine

func Map(f func(int) bool, a []int) []bool {
	lenA := 0
	for range a {
		lenA++
	}
	my := make([]bool, lenA)
	for i, v := range a {
		my[i] = f(v)
	}
	return my
}
