package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args
	len := 0
	for i := range arguments {

		len = i + 1
	}

	for s := len - 1; s > 0; s-- {

		runes := []rune(arguments[s])
		for _, k := range runes {

			z01.PrintRune(k)

		}
		z01.PrintRune('\n')
	}
}
