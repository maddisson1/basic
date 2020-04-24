package piscine

func IsPrintable(str string) bool {

	if str == "" {

		return false

	}
	var e bool
	str5 := []rune(str)
	for _, letter := range str5 {

		if letter >= ' ' && letter <= '~' {

			e = true

		} else {

			return false
		}

	}
	return e

}
