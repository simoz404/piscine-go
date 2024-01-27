package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		PrintNbr(223372036854775808)
	} else if n >= 0 && n <= 9 {
		z01.PrintRune('0' + rune(n))
	} else if n > 9 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	} else {
		z01.PrintRune('-')
		PrintNbr(-n)
	}
}
