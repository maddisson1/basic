package piscine

func Any(f func(string) bool, arr []string) bool {

	for _, c := range arr {

		if f(c) {
			return true
		}

	}
	return false
}
