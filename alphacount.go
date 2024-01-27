package piscine

func AlphaCount(s string) int {
	i := 0
	for _, value := range s {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' {
			i++
		}
	}
	return i
}
