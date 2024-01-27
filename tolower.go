package piscine

func ToLower(s string) string {
	lenS := len(s)
	str := []rune(s)
	for i := 0; i < lenS; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
