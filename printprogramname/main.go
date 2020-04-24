package main

import "github.com/01-edu/z01"
import "os"

func main() {

	arguments := os.Args[0]

	for _, letter := range arguments {

		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
