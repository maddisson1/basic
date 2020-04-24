package piscine

func IterativePower(nb int, power int) int {

	if power < 0 {

		return 0

	}
	result := 1
	for i := power; i > 0; i-- {

		result = result * nb

	}
	return result
}
