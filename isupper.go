package piscine

func IsUpper(s string) bool {
	lenS := len(s)
	str := []rune(s)
	for i := 0; i < lenS; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i == lenS-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
