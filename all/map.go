package piscine

func Map(f func(int) bool, arr []int) []bool {

	len := 0
	for i := range arr {
		len = i + 1
	}
	array_new := make([]bool, len)
	for i := range array_new {

		array_new[i] = f(arr[i])
	}
	return array_new
}
