package piscine

func len(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func Compare(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	minlen := lenA
	if lenA > lenB {
		minlen = lenB
	}
	for i := 0; i < minlen; i++ {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			} else {
				return 1
			}
		}
	}
	if lenA == lenB {
		return 0
	} else if lenA < lenB {
		return -1
	} else {
		return 1
	}
}
