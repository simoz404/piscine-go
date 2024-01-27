package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb > 0 && nb < 25 {
		result = nb * RecursiveFactorial(nb-1)
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
	return result
}
