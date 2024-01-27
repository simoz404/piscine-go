package piscine

func BasicAtoi(s string) int {
	in := 0
	for _, value := range s {
		if value >= '0' && value <= '9' {
			in = in*10 + int(value-'0')
		}
	}
	return in
}
