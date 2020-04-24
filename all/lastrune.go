package piscine

func LastRune(s string) rune {

	sentence := []rune(s)

	l := 0
	for index := range sentence {

		l = index

	}
	return sentence[l]
}
