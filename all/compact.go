package piscine

func Compact(ptr *[]string) int {

	len_res := 0
	for _, char := range *ptr {

		if char != "" {

			len_res++
		}
	}
	array := make([]string, len_res)
	c := 0
	for _, char := range *ptr {

		if char != "" {

			array[c] = char
			c++
		}
	}
	*ptr = array
	return len_res
}
