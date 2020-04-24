package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func set(ptr *point, x_changed int, y_changed int) {
	ptr.x = x_changed //42
	ptr.y = y_changed //21
}

func printStr(x_s string, x int, y_s string, y int) {

	for _, char := range x_s {

		z01.PrintRune(char)

	}
	if x > 9 {

		z01.PrintRune(gr(x / 10))
		z01.PrintRune(gr(x % 10))
	}

	for _, char := range y_s {

		z01.PrintRune(char)

	}
	z01.PrintRune(gr(y / 10))
	z01.PrintRune(gr(y % 10))

}

func gr(i int) rune {

	n := 1
	s := '1'

	for n < i {

		n++
		s++
	}
	return s

}

func main() {
	points := point{}

	set(&points, 42, 21)

	printStr("x = ", points.x, ", y = ", points.y)
	z01.PrintRune('\n')
}
