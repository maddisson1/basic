package piscine

func Index(s string, toFind string) int {

	str := []rune(s)
	sub_str := []rune(toFind)
	str_len := 0
	sub_len := 0

	for index := range str {

		index = index
		str_len++

	}
	for index := range sub_str {

		index = index
		sub_len++

	}
	if sub_len > str_len {

		return 0

	}
	counter := 0

	for i := 0; i < str_len-sub_len+1; i++ {

		if str[i] == sub_str[0] {

			a := i

			for k := 0; k < sub_len; k++ {

				if str[a] == sub_str[k] {

					counter++
				}
				a++
			}
		}
		if counter == sub_len {

			return i
		}
	}
	return -1
}
