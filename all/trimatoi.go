package piscine

func TrimAtoi(s string) int {

	len := 0
	runes := []rune(s)

	for range s {

		len++

	}
	num := 0
	nId := len + 1
	fn := len + 1
	countNeg := 0
	for i := 0; i < len; i++ {

		if runes[i] == '-' {
			nId = i
			countNeg++
		}
		if runes[i] >= 48 && runes[i] <= 57 {

			if fn > len {

				fn = i
			}
			num *= 10
			num += int(runes[i] - '0')

		}
	}
	if countNeg > 0 && fn > nId {
		return -num
	}
	return num
}
