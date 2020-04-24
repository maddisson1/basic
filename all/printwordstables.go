package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	for _, s := range table {

		runes := []rune(s)
		for _, s1 := range runes {

			z01.PrintRune(s1)

		}
		z01.PrintRune('\n')
	}
}
