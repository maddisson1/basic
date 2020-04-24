package main

import (
	"github.com/01-edu/z01"
	"os"
)

func pStr(str string) {

	arrayStr := []rune(str)
	len := 0

	for i := range arrayStr {
		len = i + 1
	}

	for i := 0; i < len; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {

	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	args := os.Args[1:]
	lenghtOfArg := 0
	for i := range args {

		lenghtOfArg = i + 1

	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lenghtOfArg) {
		pStr(EvenMsg)
	} else {
		pStr(OddMsg)
	}
}
