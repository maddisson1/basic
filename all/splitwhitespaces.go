package piscine

func SplitWhiteSpaces(str string) []string {

	s1 := ""
	len := 0
	flag := false
	for _, s := range str {

		if (s == ' ' || s == '\n' || s == '\t') && !flag {

			flag = true
			len++

		} else {

			flag = false
		}
	}
	len++
	array := make([]string, len)
	count := 0
	for _, s := range str {

		if s == ' ' || s == '\n' || s == '\t' {
			length := 0
			for i := range s1 {

				length = i + 1

			}
			if length == 0 {

				continue
			}
			array[count] = s1
			count++
			s1 = ""
			continue
		}
		s1 += string(s)
	}
	length := 0
	for i := range s1 {

		length = i + 1

	}
	if length != 0 {

		array[count] = s1

	}
	return array
}
