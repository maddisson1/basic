package piscine

func AlphaCount(str string) int {

	counter := 0
	for _, letter := range str {

		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {

			counter++

		}
	}
	return counter
}
