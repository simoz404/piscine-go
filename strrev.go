package piscine

func StrRev(s string) string {
	len := 0
	for range s {
		len++
	}
	str := []rune(s)
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		c := str[i]
		str[i] = str[j]
		str[j] = c
	}
	return string(str)
}
