package piscine

func ToUpper(s string) string {

	str := []rune(s)
	for index, letter := range str {

		if letter >= 'a' && letter <= 'z' {

			str[index] = str[index] - 32

		}
	}
	upper := string(str)
	return upper
}
