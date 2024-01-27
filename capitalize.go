package piscine

func Capitalize(s string) string {
	str := []rune(s)
	lenS := len(s)
	is := true
	for i := 0; i < lenS; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	for i := 0; i < lenS; i++ {
		if ((str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9')) && is == true {
			if str[i] >= 'a' && str[i] <= 'z' {
				str[i] = str[i] - 32
				is = false
			} else if str[i] >= '0' && str[i] <= '9' {
				is = false
			}
		} else if !((str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9')) {
			is = true
		}
	}
	return string(str)
}
