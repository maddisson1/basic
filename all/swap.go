package piscine

func Swap(a *int, b *int) {

	r := *a
	t := *b
	*a = t
	*b = r

}
