package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 20 {

		return 0

	}
	result := 1
	for nb != 0 {

		result = result * nb
		nb--

	}
	return result
}
