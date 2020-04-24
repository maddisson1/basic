package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	}
	i := nb + 1
	for {
		if IsPrime(i) == true {
			return i
		}
		i++
	}
}
