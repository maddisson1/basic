package piscine

func Abort(a, b, c, d, e int) int {

	array := []int{a, b, c, d, e}
	len := 0
	//min := 0

	for i := range array {

		len = i + 1

	}
	for k := 0; k < len-1; k++ {

		for j := 0; j < len-k-1; j++ {

			if array[j] < array[j+1] {

				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array[2]
}
