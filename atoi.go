package piscine

func Atoi(s string) int {
	out := 0
	sign := 1
	for index, value := range s {
		if index == 0 && (value == '-' || value == '+') {
			if value == '-' {
				sign = -1
			}
		} else if value >= '0' && value <= '9' {
			out = out*10 + int(value-'0')
		} else {
			return 0
		}
	}
	return out * sign
}
