package piscine

func TrimAtoi(s string) int {
	lenS := len(s)
	str := []rune(s)
	var str2 []rune
	for i := 0; i < lenS; i++ {
		if str[i] == '-' || (str[i] >= '0' && str[i] <= '9') {
			str2 = append(str2, []rune{str[i]}...)
		}
	}
	sign := 1
	result := string(str2)
	if len(result) > 0 && str2[0] == '-' {
		sign = -1
	}
	out := 0
	for _, value := range result {
		if value >= '0' && value <= '9' {
			out = out*10 + int(value-'0')
		}
	}
	return out * sign
}
