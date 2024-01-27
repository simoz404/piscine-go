package piscine

func IsLower(s string) bool {
	lenS := len(s)
	str := []rune(s)
	for i := 0; i < lenS; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			if i == lenS-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
