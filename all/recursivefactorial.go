package piscine

func RecursiveFactorial(nb int) int {

	result := 1

	if nb < 0 || nb > 20 {

		return 0

	} else if nb == 1 {

		return 1

	} else if nb > 1 {

		result = nb * RecursiveFactorial(nb-1)

	}
	return result
}
