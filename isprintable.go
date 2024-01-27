package piscine

func IsPrintable(s string) bool {
	for _, value := range s {
		if value < ' ' || value > '~' {
			return false
		}
	}
	return true
}
