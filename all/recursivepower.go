package piscine

func RecursivePower(nb int, power int) int {

	if power < 0 || power > 20 {

		return 0

	} else if power == 0 { // if power equals to 1, the same number will be returned

		return 1

	}
	result := nb * RecursivePower(nb, power-1)
	return result
}
