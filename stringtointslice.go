package piscine

func StringToIntSlice(str string) []int {
	var n []int
	for _, v := range str {
		n = append(n, int(v))
	}
	return n
}
