package piscine

func Abort(a, b, c, d, e int) int {
	nb := []int{a, b, c, d, e}
	lenA := 0
	for range nb {
		lenA++
	}
	for range nb {
		for i := 0; i < lenA-1; i++ {
			if nb[i] > nb[i+1] {
				nb[i], nb[i+1] = nb[i+1], nb[i]
			}
		}
	}
	return nb[2]
}
