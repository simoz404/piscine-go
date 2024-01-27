package piscine

func Split(s, sep string) []string {
	var s1 string
	var s2 []string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if s != "" {
				s2 = append(s2, s1)
				s1 = ""
			}
			i += len(sep) - 1
		} else {
			s1 += string(s[i])
		}
	}
	if s != "" {
		s2 = append(s2, s1)
	}
	return s2
}
