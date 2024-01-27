package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var slice []int
	if n == 0 {
		slice = append([]int{0}, slice...)
	}
	for n != 0 {
		num := n % 10
		slice = append([]int{num}, slice...)
		n = n / 10
	}
	lenS := 0
	for range slice {
		lenS++
	}
	for range slice {
		for i := 0; i < lenS-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	for j := 0; j < lenS; j++ {
		z01.PrintRune(rune(slice[j] + '0'))
	}
}
