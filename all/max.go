package piscine

func Max(arr []int) int {

	len := 0
	for i := range arr {

		len = i + 1

	}
	for j := 0; j < len-1; j++ {

		for k := 0; k < len-j-1; k++ {

			if arr[k] > arr[k+1] {

				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
	return arr[len-1]
}
