package piscine

func Capitalize(s string) string {

	res := ""
	n := 0

	for i := range s {

		n = i + 1
	}

	if s[0] >= 'a' && s[0] <= 'z' {

		res = res + string(rune(s[0]-32))

	} else {

		res = res + string(rune(s[0]))
	}
	for i := 1; i < n; i++ {

		if (s[i] >= 'a' && s[i] <= 'z') && (!isLetter(rune(s[i] - 1))) {

			res = res + string(rune(s[i]-32))
			continue
		}
		if (s[i] >= 'A' && s[i] <= 'Z') && (!isLetter(rune(s[i] - 1))) {

			res = res + string(rune(s[i]+32))
			continue
		}
		res = res + string(s[i])
	}
	return res
}

func isLetter(s rune) bool {

	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}
