package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	lenS := 0
	for range a {
		lenS++
	}
	var s string
	for i := 0; i < lenS; i++ {
		s += a[i]
		if i != lenS-1 {
			s += "\n"
		}
	}
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
