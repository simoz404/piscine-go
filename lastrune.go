package piscine

func LastRune(s string) rune {
	len := 0
	for range s {
		len++
	}
	str := []rune(s)
	return str[len-1]
}
