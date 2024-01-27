package main

import (
	"os"
)

func main() {
	lenA := len(os.Args)
	if lenA == 4 {
		n1 := 0
		n2 := 0
		sign := 1
		sign2 := 1
		const maxInt = 1<<63 - 1
		for i, v := range os.Args[1] {
			if i == 0 && (v == '-') {
				sign = -1
			} else if v >= '0' && v <= '9' {
				if n1 > (maxInt-int(v+'0'))/10 {
					return
				}
				n1 = n1*10 + int(v-'0')
			} else {
				return
			}
		}
		n1 *= sign
		for i, v := range os.Args[3] {
			if i == 0 && (v == '-') {
				sign2 = -1
			} else if v >= '0' && v <= '9' {
				if n1 > (maxInt-int(v+'0'))/10 {
					return
				}
				n2 = n2*10 + int(v-'0')
			} else {
				return
			}
		}
		is := 0
		n2 *= sign2
		if os.Args[2] == "-" || os.Args[2] == "minus" {
			var s []rune
			n := n1 - n2
			if n < 0 {
				is = 1
				n *= -1
			}
			if n == 0 {
				s = append([]rune{rune(0 + '0')}, s...)
			}
			for n > 0 {
				num := n % 10
				s = append([]rune{rune(num + '0')}, s...)
				n = n / 10
			}
			if is == 1 {
				s = append([]rune{'-'}, s...)
			}
			os.Stdout.Write([]byte(string(s) + "\n"))
		} else if os.Args[2] == "+" || os.Args[2] == "plus" {
			var s []rune
			n := n1 + n2
			if n < 0 {
				is = 1
				n *= -1
			}
			if n == 0 {
				s = append([]rune{rune(0 + '0')}, s...)
			}
			for n > 0 {
				num := n % 10
				s = append([]rune{rune(num + '0')}, s...)
				n = n / 10
			}
			if is == 1 {
				s = append([]rune{'-'}, s...)
			}
			os.Stdout.Write([]byte(string(s) + "\n"))
		} else if os.Args[2] == "*" || os.Args[2] == "times" {
			var s []rune
			n := n1 * n2
			if n < 0 {
				is = 1
				n *= -1
			}
			if n == 0 {
				s = append([]rune{rune(0 + '0')}, s...)
			}
			for n > 0 {
				num := n % 10
				s = append([]rune{rune(num + '0')}, s...)
				n = n / 10
			}
			if is == 1 {
				s = append([]rune{'-'}, s...)
			}
			os.Stdout.Write([]byte(string(s) + "\n"))
		} else if os.Args[2] == "/" || os.Args[2] == "dividedby" {
			if n2 == 0 {
				s := "No division by 0"
				os.Stdout.Write([]byte(s + "\n"))
			} else {
				var s []rune
				n := n1 / n2
				if n < 0 {
					is = 1
					n *= -1
				}
				if n == 0 {
					s = append([]rune{rune(0 + '0')}, s...)
				}
				for n > 0 {
					num := n % 10
					s = append([]rune{rune(num + '0')}, s...)
					n = n / 10
				}
				if is == 1 {
					s = append([]rune{'-'}, s...)
				}
				os.Stdout.Write([]byte(string(s) + "\n"))
			}
		} else if os.Args[2] == "%" || os.Args[2] == "modulo" {
			if n2 == 0 {
				s := "No modulo by 0"
				os.Stdout.Write([]byte(s + "\n"))
			} else {
				var s []rune
				n := n1 % n2
				if n < 0 {
					is = 1
					n *= -1
				}
				if n == 0 {
					s = append([]rune{rune(0 + '0')}, s...)
				}
				for n > 0 {
					num := n % 10
					s = append([]rune{rune(num + '0')}, s...)
					n = n / 10
				}
				if is == 1 {
					s = append([]rune{'-'}, s...)
				}
				os.Stdout.Write([]byte(string(s) + "\n"))
			}
		}
	}
	return
}
