package piscine

func JumpOver(str string) string {
	var s []rune
	ss := []rune(str)
	for i := 2; i < len(str); i += 3 {
		s = append(s, ss[i])
	}
	s = append(s, '\n')
	return string(s)
}
