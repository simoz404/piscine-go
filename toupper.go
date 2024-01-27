package piscine

func ToUpper(s string) string {
	lenS := len(s)
	str := []rune(s)
	for i := 0; i < lenS; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
