package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	arr := make([]byte, 1024)
	exit := 1
	if len(arg) == 0 {
		content, err := os.Stdin.Read(arr)
		if err == nil {
			PrintStr(string(arr[:content]))
		}
	}
	for i := 0; i < len(arg); i++ {
		content, err := ioutil.ReadFile(arg[i])
		if err != nil {
			PrintStr("ERROR: open " + arg[i] + ": no such file or directory\n")
			os.Exit(exit)
			return
		}
		PrintStr(string(content))
	}
}

func PrintStr(v string) {
	for i := 0; i < len(v); i++ {
		z01.PrintRune(rune(v[i]))
	}
}
