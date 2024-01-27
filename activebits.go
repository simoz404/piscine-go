package piscine

func ActiveBits(n int) int {
	var nb []int
	for n >= 1 {
		if n%2 == 1 {
			nb = append(nb, n%2)
		}
		n /= 2
	}
	lenN := 0
	for range nb {
		lenN++
	}
	return lenN
}
