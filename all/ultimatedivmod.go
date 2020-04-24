package piscine

func UltimateDivMod(a *int, b *int) {

	k := *a
	l := *b
	*a = k / l
	*b = k % l

}
