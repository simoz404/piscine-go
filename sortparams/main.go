package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenS := len(os.Args)
	for range os.Args {
		for i := 1; i < lenS-1; i++ {
			if os.Args[i] > os.Args[i+1] {
				c := os.Args[i]
				os.Args[i] = os.Args[i+1]
				os.Args[i+1] = c
			}
		}
	}
	for i := 1; i < lenS; i++ {
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
