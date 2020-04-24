package piscine

func IsLower(str string) bool {

	if str == "" {

		return false

	}
	var b bool
	str2 := []rune(str)
	for _, letter := range str2 {

		if letter >= 'a' && letter <= 'z' {

			b = true

		} else {

			return false
		}
	}
	return b
}
