package piscine

func Prime(nb int) bool {
	if nb == 2 {
		return true
	} else if nb%2 == 0 || nb <= 1 {
		return false
	} else {
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	} else if Prime(nb) == true {
		return nb
	}
	return FindNextPrime(nb + 1)
}
