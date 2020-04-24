package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n < 0 {

		return

	}
	if n == 0 {

		z01.PrintRune('0')

	}
	var array [10]int // creating an array to append
	for n != 0 {
		array[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {

		for array[i] > 0 {

			z01.PrintRune(rune(i) + '0')
			array[i]--
		}
	}
}
