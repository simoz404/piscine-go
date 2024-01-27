package piscine

func ReverseMenuIndex(menu []string) []string {
	var s []string
	lenS := 0
	for range menu {
		lenS++
	}
	for i := lenS - 1; i >= 0; i-- {
		s = append(s, menu[i])
	}
	return s
}
