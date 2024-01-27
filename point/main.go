package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printInt(n int) {
	nb := '0'
	for i := 0; i < n/10; i++ {
		nb++
	}
	z01.PrintRune(nb)
	nb = '0'
	for i := 0; i < n%10; i++ {
		nb++
	}
	z01.PrintRune(nb)
}

func main() {
	points := &point{}

	setPoint(points)
	s := "x = i, y = j"
	for _, v := range s {
		if v == 'i' {
			printInt(points.x)
		} else if v == 'j' {
			printInt(points.y)
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
