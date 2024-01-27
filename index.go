package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenT := len(toFind)

	if lenS == 0 || lenT == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenT; i++ {
		for j := 0; j < lenT; j++ {
			if s[i+j] != toFind[j] {
				break
			}
			if j == lenT-1 {
				return i
			}
		}
	}

	return -1
}
