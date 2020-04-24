package piscine

func Sqrt(nb int) int {

	if nb == 1 {

		return 1

	}

	sqrtn := 0
	for i := 0; i <= nb/2; i++ {

		if i*i == nb {

			sqrtn = i

		}
	}
	return sqrtn
}
