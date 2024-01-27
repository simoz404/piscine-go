package piscine

func ShoppingListSort(slice []string) []string {
	lenS := 0
	for range slice {
		lenS++
	}
	for range slice {
		for i := 0; i < lenS-1; i++ {
			if len(slice[i]) > len(slice[i+1]) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	return slice
}
