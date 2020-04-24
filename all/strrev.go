package piscine

func StrRev(s string) string {

	ordstr := []rune(s)
	revstr := []rune(s)

	var r int
	for i := range ordstr {

		r = i

	}

	for i := range revstr {

		revstr[i] = ordstr[r]
		r = r - 1

	}
	return string(revstr)
}
