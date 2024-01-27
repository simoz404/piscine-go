package piscine

func LoafOfBread(str string) string {
	var result []rune

	if str == "" {
		result = append(result, '\n')
		return string(result)
	}

	if len(str) < 5 {
		result = append(result, []rune("Invalid Output\n")...)
		return string(result)
	}

	j := 0
	for i := 0; i < len(str); i++ {
		if j < 5 && rune(str[i]) == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			result = append(result, ' ')
			j = 0
			continue
		}
		result = append(result, rune(str[i]))
		j++
	}

	result = append(result, '\n')
	return string(result)
}
