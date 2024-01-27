package piscine

func Unmatch(a []int) int {
	var qount int
	for _, el := range a {
		qount = 0
		for _, v := range a {
			if v == el {
				qount++
			}
		}
		if qount%2 != 0 {
			return el
		}
	}
	return -1
}
