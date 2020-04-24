package piscine

func ToLower(s string) string {

	str := []rune(s)
	for index, letter := range str {

		if letter >= 'A' && letter <= 'Z' {

			str[index] = str[index] + 32

		}
	}
	lower := string(str)
	return lower
}
