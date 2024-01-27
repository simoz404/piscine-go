package piscine

func IsAlpha(s string) bool {
	lenS := len(s)
	str := []rune(s)
	for i := 0; i < lenS; i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9') {
			if i == lenS-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
