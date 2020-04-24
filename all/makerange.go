package piscine

func MakeRange(min, max int) []int {

	var answer []int

	if min < max {
		answer = make([]int, max-min)
		for i := 0; i < max-min; i++ {
			answer[i] = min + i
		}
	}
	return answer
}
