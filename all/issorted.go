package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {

	len := 0
	for i := range tab {
		len = i + 1
	}
	l := false
	r := false
	for i := 0; i < len-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			l = true
		} else if f(tab[i], tab[i+1]) < 0 {
			r = true
		}
	}
	if l && r {
		return false
	} else {
		return true
	}
}
