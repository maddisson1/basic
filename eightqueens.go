package main

import (
	"github.com/01-edu/z01"
)

func putQueen(positions []int, row int, size int) {
	if row == size {
		printAns(positions, size)
	} else {
		for col := 0; col < size; col++ {
			if checkPlace(positions, row, col) {
				positions[row] = col
				putQueen(positions, row+1, size)
			}
		}
	}
}
func checkPlace(positions []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if positions[i] == col || positions[i]-i == col-row || positions[i]+i == col+row {
			return false
		}
	}
	return true
}
func printAns(positions []int, size int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if positions[row] == col {
				z01.PrintRune(rune(col + 1 + '0'))
			}
		}
	}
	z01.PrintRune('\n')
}

func EightQueens() {
	size := 8
	positions := []int{-1, -1, -1, -1, -1, -1, -1, -1}
	putQueen(positions, 0, size)
}
