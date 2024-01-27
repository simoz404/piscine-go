package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	lenA := 0
	for range a {
		lenA++
	}
	for i := 0; i < lenA-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			if (i+2 <= lenA-1) && f(a[i+1], a[i+2]) < 0 {
				return false
			}
		} else if f(a[i], a[i+1]) < 0 {
			if (i+2 <= lenA-1) && f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}
	return true
}
