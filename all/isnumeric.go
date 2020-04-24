package piscine

func IsNumeric(str string) bool {

	if str == "" {

		return false

	}
	var c bool
	str4 := []rune(str)
	for _, letter := range str4 {

		if letter >= '0' && letter <= '9' {

			c = true

		} else {

			return false
		}
	}
	return c
}
