package piscine

func IsUpper(str string) bool {

	if str == "" {

		return false

	}
	var a bool
	str1 := []rune(str)
	for _, letter := range str1 {

		if letter >= 'A' && letter <= 'Z' {

			a = true

		} else {

			return false

		}
	}
	return a
}
