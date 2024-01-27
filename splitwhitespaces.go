package piscine

func SplitWhiteSpaces(s string) []string {
	var str string
	var result []string
	for _, v := range s {
		if v == '\n' || v == ' ' || v == '\t' {
			if str != "" {
				result = append(result, str)
				str = ""
			}
		} else {
			str += string(v)
		}
	}
	if str != "" {
		result = append(result, str)
	}
	return result
}
