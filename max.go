package piscine

func Max(a []int) int {
	lenA := 0
	for range a {
		lenA++
	}
	for range a {
		for i := 0; i < lenA-1; i++ {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a[0]
}
