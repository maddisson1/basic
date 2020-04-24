package piscine

func StrLen(str string) int {

	a := []rune(str)

	var r int
	for i := range a {

		r = i

	}
	return r + 1
}
