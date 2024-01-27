package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	lenS := max - min
	s := make([]int, lenS)
	for i := 0; i < lenS; i++ {
		s[i] = min + i
	}
	return s
}
