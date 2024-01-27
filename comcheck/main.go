package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args
	for i := 0; i < len(s); i++ {
		if s[i] == "01" || s[i] == "galaxy" || s[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
