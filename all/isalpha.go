package piscine

func IsAlpha(str string) bool {

	if str == "" {

		return false

	}
	str3 := []rune(str)
	var b bool
	for _, letter := range str3 {

		if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' || letter >= '0' && letter <= '9' {

			b = true
		} else {

			return false

		}
	}
	return b
}
