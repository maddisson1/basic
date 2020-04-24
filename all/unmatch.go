package piscine

func Unmatch(arr []int) int {

	len := 0
	count := 0
	for i := range arr {

		len = i + 1

	}
	for i := 0; i < len; i++ {

		for j := 0; j < len; j++ {

			if arr[i] == arr[j] {

				count++
			}
		}
		if count%2 == 1 {

			return arr[i]
		}
	}
	return -1
}
