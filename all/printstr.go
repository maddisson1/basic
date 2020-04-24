package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	a := []rune(str)

	for i := range string(a) {

		z01.PrintRune(a[i])

	}
}
