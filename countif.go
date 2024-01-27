package piscine

func CountIf(f func(string) bool, tab []string) int {
	i := 0
	for _, v := range tab {
		if f(v) == true {
			i++
		}
	}
	return i
}
