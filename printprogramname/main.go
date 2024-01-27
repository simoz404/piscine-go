package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	first := os.Args[0]
	for _, v := range first {
		if v == '.' || v == '/' {
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
