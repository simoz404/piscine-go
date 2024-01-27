package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	i := 1
	if len(s) == 1 || (s[1] == "--upper" && len(s) == 1) {
		return
	}
	if s[1] == "--upper" {
		i = 2
	}
	var out rune
	for ; i < len(s); i++ {
		sint := 0
		for _, let := range s[i] {
			if let >= '0' && let <= '9' {
				sint = sint*10 + int(rune(let)-'0')
			}
		}
		if sint <= 26 && sint >= 1 {
			if s[1] == "--upper" {
				out = rune(sint + 64)
			} else {
				out = rune(sint + 96)
			}
			z01.PrintRune(out)
		} else {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
}
