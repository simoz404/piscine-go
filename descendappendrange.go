package piscine

func DescendAppendRange(max, min int) []int {
	n := []int{}
	if max < min {
		return n
	}
	for i := max; i > min; i-- {
		n = append(n, i)
	}
	return n
}
