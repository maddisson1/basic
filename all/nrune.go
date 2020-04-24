package piscine

func NRune(s string, n int) rune {

	sentence := []rune(s)

	for index, letter := range sentence {

		if index == n-1 {

			return letter

		}
	}
	return 0
}
