package piscine

func SortIntegerTable(table []int) {
	lentable := 0
	for range table {
		lentable++
	}
	for range table {
		for i := 0; i < lentable-1; i++ {
			if table[i] > table[i+1] {
				c := table[i]
				table[i] = table[i+1]
				table[i+1] = c
			}
		}
	}
}
