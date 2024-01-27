package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenS := len(os.Args)
	for i := lenS - 1; i >= 1; i-- {
		s := os.Args[i]
		for _, v := range s {
			if v == '.' || v == '/' {
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
