package piscine

func NRune(s string, n int) rune {
	len := 0
	for range s {
		len++
	}
	if n > 0 && n <= len {
		str := []rune(s)
		return str[n-1]
	}
	return 0
}
